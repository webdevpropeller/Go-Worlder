package storage

import (
	"math/rand"
	"mime/multipart"
	"time"

	"context"
	"io"
	"os"
	"strings"

	"cloud.google.com/go/storage"
	log "github.com/sirupsen/logrus"
)

const (
	fileNameLength  int = 32
	fileContentType int = 1

	developmentBucket string = "wolonote-development-data-asia-northeast1"
	productionBucket  string = ""

	BrandPath      string = "assets/images/brands/"
	BackGroundPath string = "assets/images/users/backgrounds/"
	IconPath       string = "assets/images/users/icons/"
	ProductPath    string = "assets/images/products/"
)

var (
	bucketName string
	mediaLink  string
	rs1Letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

// UploadFileForCreate ...
func UploadFileForCreate(fh *multipart.FileHeader, objectPath string) error {
	if fh == nil {
		return nil
	}

	fileName := generateFileName(fh)
	// Get file data
	file, err := fh.Open()
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	env := os.Getenv("GO_ENV")
	// Set object name
	objectName := objectPath + fileName
	if env == "local" {
		// Write object to container
		err = writeObjectToLocal(objectName, file, fh)
		if err != nil {
			return err
		}
	} else if env == "development" || env == "production" {
		setBucketName(env)
		// Create a client
		ctx, client, err := createClient()
		if err != nil {
			return err
		}

		// Write object to container and get media link
		err = writeObjectToStorage(ctx, client, objectName, file, fh)
		if err != nil {
			return err
		}
	}

	return nil
}

// UploadFileForUpdate ...
func UploadFileForUpdate(fh *multipart.FileHeader, objectPath string) error {
	if fh == nil {
		return nil
	}

	fileName := generateFileName(fh)
	// Get file data
	file, err := fh.Open()
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	env := os.Getenv("GO_ENV")
	// Set object name
	objectName := objectPath + fileName
	if env == "local" {
		// Check file exists
		_, err := os.Stat(objectPath + fh.Filename)
		if err == nil {
			log.Println(err)
			return err
		}

		// Write object to container
		err = writeObjectToLocal(objectName, file, fh)
		if err != nil {
			return err
		}
	} else if env == "development" || env == "production" {
		setBucketName(env)
		// Create a client
		ctx, client, err := createClient()
		if err != nil {
			return err
		}

		// Check file exists
		rc, err := client.Bucket(bucketName).Object(objectPath + fh.Filename).NewReader(ctx)
		if err == nil {
			log.Println(err)
			return err
		}
		// Complete to read object
		if err := rc.Close(); err != nil {
			log.Println(err)
			return err
		}

		// Write object to container and get media link
		err = writeObjectToStorage(ctx, client, objectName, file, fh)
		if err != nil {
			return err
		}
	}

	return nil
}

func generateFileName(fh *multipart.FileHeader) string {
	contentType := fh.Header.Get("Content-Type")
	sliceContentType := strings.Split(contentType, "/")
	fileName := generateRandomString(fileNameLength) + "." + sliceContentType[fileContentType]
	return fileName
}

func writeObjectToLocal(objectName string, file multipart.File, fh *multipart.FileHeader) error {
	// Create container
	dstFile, err := os.Create(objectName)
	if err != nil {
		log.Println(err)
		return err
	}
	defer dstFile.Close()

	// Write object to the container
	_, err = io.Copy(dstFile, file)
	if err != nil {
		log.Println(err)
		return err
	}

	mediaLink = objectName
	fh.Filename = mediaLink

	return nil
}

func createClient() (context.Context, *storage.Client, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}
	return ctx, client, nil
}

func setBucketName(env string) {
	if env == "development" {
		bucketName = developmentBucket
	} else if env == "production" {
		bucketName = productionBucket
	}
}

func writeObjectToStorage(ctx context.Context, client *storage.Client, objectName string, file multipart.File, fh *multipart.FileHeader) error {
	// Create write container
	wc := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)

	// Write object to the container
	if _, err := io.Copy(wc, file); err != nil {
		log.Println(err)
		return err
	}

	// Complete to write object
	if err := wc.Close(); err != nil {
		log.Println(err)
		return err
	}

	fh.Filename = wc.Attrs().MediaLink
	return nil
}

func generateRandomString(length int) (str string) {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)
	for i := range b {
		b[i] = rs1Letters[rand.Intn(len(rs1Letters))]
	}
	str = string(b)
	return
}

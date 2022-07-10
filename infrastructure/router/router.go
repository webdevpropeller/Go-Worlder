package router

import (
	"fmt"
	"go_worlder_system/infrastructure/cache"
	"go_worlder_system/infrastructure/sqlhandler"
	"go_worlder_system/interface/controller"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stripe/stripe-go"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Start ...
func Start() {
	sqlHandler := sqlhandler.NewSQLHandler()
	nosqlHandler := sqlhandler.NewRedisClient()
	cache := cache.New()
	// Create handler
	userController := controller.NewUserController(sqlHandler)
	profileController := controller.NewProfileController(sqlHandler)
	brandController := controller.NewBrandController(sqlHandler)
	brandLikeController := controller.NewBrandLikeController(sqlHandler)
	productController := controller.NewProductController(sqlHandler)
	projectController := controller.NewProjectController(sqlHandler)
	searchController := controller.NewSearchController(sqlHandler)
	accountController := controller.NewAccountController(sqlHandler)
	orderController := controller.NewOrderController(sqlHandler)
	transactionController := controller.NewTransactionController(sqlHandler)
	payeeController := controller.NewPayeeController(sqlHandler)
	inventoryController := controller.NewInventoryController(sqlHandler)
	chatController := controller.NewChatController(sqlHandler, nosqlHandler, cache)

	e := echo.New()

	e.Use(WrapContext)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowCredentials: true,
	}))
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		path := c.Request().URL.Path
		if strings.Contains(path, "swagger") {
			return
		}
		fmt.Printf("Request Body: %v\n", string(reqBody))
		fmt.Printf("Response Body: %v\n", string(resBody))
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: logFormat(),
		Output: os.Stdout,
	}))
	e.Use(middleware.Recover())

	// Setting settlement
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	// routing
	api := e.Group("/api")
	api.GET("/", c(userController.Home), logout)

	// user authentication routing
	api.GET("/signup", c(userController.New), logout)
	api.POST("/signup", c(userController.SignUp), logout)
	api.POST("/activate", c(userController.Activate), logout)
	api.GET("/signin", c(userController.NewSignin), logout)
	api.POST("/signin", c(userController.SignIn), logout)
	api.POST("/password/forgot", c(userController.ForgotPassword), logout)
	api.POST("/password/reset", c(userController.ResetPassword), logout)
	api.POST("/signout", c(userController.Signout), login)

	// Profile routing
	profile := api.Group("/profile")
	profile.GET("/:id", c(profileController.Show))
	profile.GET("/new", c(profileController.New), login)
	profile.GET("/edit", c(profileController.Edit), login)
	profile.POST("", c(profileController.Create), login)
	profile.PATCH("", c(profileController.Update), login)
	profile.GET("/brandlike", c(profileController.IndexBrandLike), login)

	brands := api.Group("/brands")
	// brands routing
	brands.GET("", c(brandController.Index), login)
	brands.GET("/:id", c(brandController.Show), login)
	brands.GET("/new", c(brandController.New), login)
	brands.GET("/:id/edit", c(brandController.Edit), login)
	brands.POST("", c(brandController.Create), login)
	brands.PATCH("/:id", c(brandController.Update), login)
	brands.DELETE("/:id", c(brandController.Delete), login)
	brands.GET("/:id/like", c(brandLikeController.Index), login)
	brands.POST("/:id/like", c(brandLikeController.Create), login)
	brands.DELETE("/:id/like", c(brandLikeController.Delete), login)

	// products routing
	products := api.Group("/products")
	products.GET("", c(productController.Index), login)
	products.GET("/:id", c(productController.Show), login)
	products.GET("/new", c(productController.New), login)
	products.GET("/:id/edit", c(productController.Edit), login)
	products.POST("", c(productController.Create), login)
	products.PATCH("/:id", c(productController.Update), login)
	products.DELETE("/:id", c(productController.Delete), login)

	// projects routing
	projects := api.Group("/projects")
	projects.GET("", c(projectController.Index), login)
	projects.GET("/:id", c(projectController.Show), login)
	projects.GET("/new", c(projectController.New), login)
	projects.GET("/:id/edit", c(projectController.Edit), login)
	projects.POST("", c(projectController.Create), login)
	projects.PATCH("/:id", c(projectController.Update), login)
	projects.DELETE("/:id", c(projectController.Delete), login)

	// search routing
	search := api.Group("/search")
	search.GET("/product", c(searchController.SearchProduct), login)

	// accounting routing
	accounting := api.Group("/accounting")
	// accounting accounts routing
	accounts := accounting.Group("/accounts")
	accounts.GET("", c(accountController.Index), login)
	accounts.GET("/:id", c(accountController.Show), login)
	accounts.GET("/new", c(accountController.New), login)
	accounts.POST("", c(accountController.Create), login)
	accounts.DELETE("/:id", c(accountController.Delete), login)
	// accounting transaction routing
	transactions := accounting.Group("/transactions")
	transactions.GET("", c(transactionController.Index), login)
	transactions.GET("/accounts/:id", c(transactionController.IndexAccount), login)
	transactions.POST("", c(transactionController.Create), login)
	transactions.GET("/new", c(transactionController.New), login)
	transactions.GET("/:id/edit", c(transactionController.Edit), login)
	transactions.PATCH("/:id", c(transactionController.Update), login)
	transactions.DELETE("/:id", c(transactionController.Delete), login)
	// transactions payees routing
	payees := transactions.Group("/payees")
	payees.GET("", c(payeeController.Index), login)
	payees.POST("", c(payeeController.Create), login)
	payees.PATCH("/:id", c(payeeController.Update), login)
	payees.DELETE("/:id", c(payeeController.Delete), login)

	// inventory routing
	inventory := api.Group("/inventory")
	// invengory list routing
	list := inventory.Group("/list", login)
	list.GET("", c(inventoryController.Index), login)
	list.GET("/:id", c(inventoryController.Show), login)
	list.GET("/:id/edit", c(inventoryController.Edit), login)
	list.PATCH("/:id", c(inventoryController.Update), login)
	// inventory receiving routing
	receiving := inventory.Group("/receiving")
	receiving.GET("", c(inventoryController.NewReceiving), login)
	receiving.POST("", c(inventoryController.CreateReceiving), login)
	// inventory shipping routing
	shipping := inventory.Group("/shipping")
	shipping.GET("", c(inventoryController.NewShipping), login)
	shipping.POST("", c(inventoryController.CreateShipping), login)
	// inventory disposal routing
	disposal := inventory.Group("/disposal")
	disposal.GET("", c(inventoryController.NewDisposal), login)
	disposal.POST("", c(inventoryController.CreateDisposal), login)
	// inventory stocktaking routing
	stocktaking := inventory.Group("/stocktaking")
	stocktaking.GET("", c(inventoryController.NewStocktaking), login)
	stocktaking.POST("", c(inventoryController.CreateStocktaking), login)

	// order routing
	orders := api.Group("/orders")
	orders.POST("", c(orderController.Create), login)

	// chat
	go chatController.BroadCast()
	chat := api.Group("/chat")
	chat.GET("/ws", c(chatController.WebSocket))
	chat.GET("/destinations", c(chatController.IndexDestination), login)
	chat.GET("/:id", c(chatController.Show), login)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func logFormat() string {
	format := strings.Replace(middleware.DefaultLoggerConfig.Format, ",", ",\n  ", -1)
	format = strings.Replace(format, "{", "{\n  ", 1)
	format = strings.Replace(format, "}}", "}\n}", 1)
	return format
}

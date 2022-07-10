package sqlhandler

import "github.com/go-redis/redis"

// RedisClient ...
type RedisClient struct {
	*redis.Client
}

// NewRedisClient ...
func NewRedisClient() *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &RedisClient{rdb}
}

// HSet ...
func (client *RedisClient) HSet(key string, field string, value interface{}) error {
	return client.Client.HSet(key, field, value).Err()
}

// HGet ...
func (client *RedisClient) HGet(key string, field string) (string, error) {
	return client.Client.HGet(key, field).Result()
}

// HGetAll ...
func (client *RedisClient) HGetAll(id string) (map[string]string, error) {
	return client.Client.HGetAll(id).Result()
}

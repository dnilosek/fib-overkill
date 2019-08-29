package database

import "github.com/go-redis/redis"

type RedisDB struct {
	Client *redis.Client
}

// Close connection
func (db *RedisDB) Close() error {
	return db.Client.Close()
}

// Set key/value pair
func (db *RedisDB) Set(key, value string) error {
	return db.Client.Set(key, value, 0).Err()
}

// HSet key/value
func (db *RedisDB) HSet(key, field, value string) error {
	return db.Client.HSet(key, field, value).Err()
}

// HGet value
func (db *RedisDB) HGet(key, field string) (string, error) {
	return db.Client.HGet(key, field).Result()
}

// HGet all
func (db *RedisDB) HGetAll(key string) (map[string]string, error) {
	return db.Client.HGetAll(key).Result()
}

// Get value from key
func (db *RedisDB) Get(key string) (string, error) {
	return db.Client.Get(key).Result()
}

// Open connection to redis db
func OpenRedis(databaseAddr string) (*RedisDB, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     databaseAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	db := RedisDB{
		Client: client,
	}
	return &db, nil
}

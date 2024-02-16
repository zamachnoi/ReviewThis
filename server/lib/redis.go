package lib

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

// redis client
var ctx = context.Background()
var RD *redis.Client

func InitRD() {
    opt, err := redis.ParseURL(os.Getenv("UPSTASH_REDIS_URL"))
    if err != nil {
        panic("failed to parse Redis URL")
    }

    RD = redis.NewClient(opt)
}

func CloseRD() (error){
    err := RD.Close()
    return err
}

func GetRD() *redis.Client {
    log.Printf("Getting Redis client")
    return RD
}

func SetCache(key string, value string, expiration time.Duration) (error) {
    log.Printf("Setting cache for key: %s", key)
    err := RD.Set(ctx, key, value, expiration).Err()
    if err != nil {
        log.Printf("Error setting cache for key: %s, error: %v", key, err)
        return err
    } else {
        log.Printf("Cache set for key: %s value: %s", key, value)
    }
    return nil
}

func GetCache(key string) (string, error) {
    log.Printf("Getting cache for key: %s", key)
    value, err := RD.Get(ctx, key).Result()
    if err == redis.Nil {
        log.Printf("Cache miss for key: %s", key)
        return "", err
    } else if err != nil {
        log.Printf("Error getting cache for key: %s, error: %v", key, err)
        return "", err
    }
    log.Printf("Cache hit for key: %s, value: %s", key, value)
    return value, nil
}

func DeleteCache(key string) error {
    log.Printf("Deleting cache for key: %s", key)
    err := RD.Del(ctx, key).Err()
    if err != nil {
        log.Printf("Error deleting cache for key: %s, error: %v", key, err)
        return err
    }
    return nil
}
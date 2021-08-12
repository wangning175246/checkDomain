package redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

var Client *redis.Client

func InitClient() error {
	// 创建 redis 客户端
	Client = redis.NewClient(&redis.Options{
		Addr:     "192.168.124.128:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 10, // 连接池的大小为 10
	})
	pong, err := Client.Ping().Result()
	fmt.Println("Redis Client: " + pong)
	return err
}

//func CreateProducers()  {
//	Client.XGroupCreate(Mode.REDIS_EXECMQ,"")
//}
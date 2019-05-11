package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var client *redis.Client

func init() {
	newRedisClient()
}

func newRedisClient() {

	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		//Password: "", // no password set
		DB: 10, // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

func redisGet(key string) *string {
	result, err := client.Get(key).Result()
	if err != nil {
		panic(err)
	}
	return &result
}

func redisSet(key string, value interface{}, duration time.Duration) {
	err := client.Set(key, value, duration).Err()
	if err != nil {
		panic(err)
	}
}

func redisDel(key string) {
	client.Del(key)
}

func main() {

	//value := *redisGet("aliorder:productlist:1229006772465588")
	//fmt.Println(value)
	//v := string("www.zoolye.com")
	//redisSet("area:zool", v, 0)
	//redisDel("area:zool")
}

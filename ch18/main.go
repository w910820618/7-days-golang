package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var (
	client *redis.Client
	ctx    context.Context
)

func InitConn() {
	option := &redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	}
	client = redis.NewClient(option)

	ctx = context.Background()

	ping, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("redis 连接失败:", err.Error())
		return
	}
	fmt.Println("redis 连接成功", ping)
}

func SetValue(key string, value interface{}) {
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func GetValue(key string) {
	//  get 一个值
	val2, err := client.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("读取：", val2)
}

func main() {
	InitConn()
	SetValue("name", "zhouyang")
	GetValue("name")
}

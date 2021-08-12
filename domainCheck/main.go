package main

import (
	"CheckDomain/Mode"
	RedCLi "CheckDomain/dal/redis"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func Consumer(ch chan string)  {
	args := redis.XReadGroupArgs{
		Group:    "group01",
		Consumer: "consumer_2",
		Streams:  []string{Mode.REDIS_EXECMQ,">"},
		Count:    1,
		Block:    0,
		NoAck:    false,
	}
	for   {
		result,err:=RedCLi.Client.XReadGroup(&args).Result()
		if err!=nil{
			fmt.Println(err)
			return
		}
		str,ok:=result[0].Messages[0].Values["domain"].(string)
		if !ok{
			fmt.Println("转换失败")
			continue
		}
		fmt.Println(str)
		ch<-str
		time.Sleep(200)
	}
}

func main()  {
	DomainListChan:=make(chan string,1024)
	err := RedCLi.InitClient()
	if err != nil {
		fmt.Printf("init redis errr", err)
		return
	}
	go Consumer(DomainListChan)
	time.Sleep(100000000000000)
}

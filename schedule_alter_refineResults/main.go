package main

import (
	"CheckDomain/dal/redis"
	"CheckDomain/schedule_alter_refineResults/schedule"
	"fmt"
	"github.com/robfig/cron"
)


func main()  {
	err:=redis.InitClient()
	if err != nil {
		fmt.Println("redis init fail",err.Error())
	}
	c := cron.New()
	c.AddFunc("0 */1 * * * ?",schedule.Schedul)
	c.Start()
	select {
	}

}

package main

import (
	"CheckDomain/Mode"
	"CheckDomain/dal/redis"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)


func main() {
	err:= redis.InitClient()
	if err!=nil{
		fmt.Printf("redis init fail",err)
		return
	}
	route := gin.Default()
	route.POST("/task",AddTask)
	err=route.Run(":8080")
	if err != nil {
		log.Fatalf("server start fail:",err)
	}
}

func AddTask(c *gin.Context) {
	var taskList []*Mode.TaskInfo
	err:=c.ShouldBindJSON(&taskList)
	if err!=nil{
		fmt.Println(err.Error())
		c.JSON(500,gin.H{
			"code":999,
			"message":err.Error(),
			"data":"",
		})
		return
	}
	fmt.Printf("%#v\n",&taskList)
	var isAddSucces bool=true
	var addFailDomain = map[string]error{}
	for _,task :=range taskList{
		data,err:=json.Marshal(task)
		if err != nil {
			fmt.Println(err)
			isAddSucces=false
			addFailDomain[task.DomainName]=err
			continue
		}
		rest:= redis.Client.HSet(Mode.REDIS_DOMAINLISK_KEY,task.DomainName,data)
		if rest.Err() !=nil{
			fmt.Println(err)
			isAddSucces=false
			addFailDomain[task.DomainName]=err
			continue
		}
	}
	if !isAddSucces{
		c.JSON(500,gin.H{
			"code":999,
			"message":addFailDomain,
			"data":"",
		})
		return
	}
	c.JSON(200,gin.H{
		"code":0,
		"message":"success",
		"data":"",
	})
}

package schedule

import (
	"CheckDomain/Mode"
	RedCLi "CheckDomain/dal/redis"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func Schedul()  {
	kes,err:=RedCLi.Client.HKeys(Mode.REDIS_DOMAINLISK_KEY).Result()
	if err!=nil{
		fmt.Println("redis get key fail:",err.Error())
	}
	for _,domain:=range kes {
		var isModifi bool =false
		data,err:=RedCLi.Client.HGet(Mode.REDIS_DOMAINLISK_KEY,domain).Result()
		if err != nil {
			fmt.Printf("redis get %s from %s fail:",domain,Mode.REDIS_DOMAINLISK_KEY)
		}
		var taskinfo=&Mode.TaskInfo{}
		err=json.Unmarshal([]byte(data),taskinfo)
		if err!=nil{
			fmt.Printf("json Unmarshal fail")
		}
		currentTime:=time.Now().Unix()

		switch taskinfo.CheckRate {
		case "10m":
			if currentTime-taskinfo.LastExecutedEvent>=Mode.M10{
				err:=Producers(domain)
				if err != nil {
					fmt.Printf("redis push %s from %s fail:%v",domain,Mode.REDIS_EXECMQ,err)
				}else{
					isModifi=true
					taskinfo.LastExecutedEvent=currentTime
				}
			}
		case "30m":
			if currentTime-taskinfo.LastExecutedEvent>=Mode.M30{
				err:=Producers(domain)
				if err != nil {
					fmt.Printf("redis push %s from %s fail:%v",domain,Mode.REDIS_EXECMQ,err)
				}else{
					isModifi=true
					taskinfo.LastExecutedEvent=currentTime
				}
			}
		case "1h":
			if currentTime-taskinfo.LastExecutedEvent>=Mode.H1{
				err:=Producers(domain)
				if err != nil {
					fmt.Printf("redis push %s from %s fail:%v",domain,Mode.REDIS_EXECMQ,err)
				}else{
					isModifi=true
					taskinfo.LastExecutedEvent=currentTime
				}
			}
		case "24h":
			if currentTime-taskinfo.LastExecutedEvent>=Mode.H24{
				err:=Producers(domain)
				if err != nil {
					fmt.Printf("redis push %s from %s fail:",domain,Mode.REDIS_EXECMQ)
				}else{
					isModifi=true
					taskinfo.LastExecutedEvent=currentTime
				}
			}
		default:
			fmt.Printf("doamin:%s schedule fail",domain)
		}
		if isModifi{
			modifiData,err:=json.Marshal(taskinfo)
			if err!=nil{
				fmt.Println("json Marshal fail ",err)
			}else {
				rest:=RedCLi.Client.HSet(Mode.REDIS_DOMAINLISK_KEY,domain,modifiData)
				if rest.Err()!=nil{
					fmt.Printf("edit modifi domain %s taskInfo fail",domain)
				}
			}
		}
	}
}

func Producers(domain string) (err error) {
	xadd_args:=redis.XAddArgs{
		Stream:       Mode.REDIS_EXECMQ,
		MaxLen:       0,
		MaxLenApprox: 0,
		ID:           "*",
		Values: map[string]interface{}{
			"domain":domain,
		},
	}
	str,err:=RedCLi.Client.XAdd(&xadd_args).Result()
	fmt.Println(str)
	fmt.Println(err)
	return
}
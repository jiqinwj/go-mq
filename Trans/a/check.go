package main

import (
	"github.com/robfig/cron/v3"
	"log"
	"go-mq/Trans"
)
var MyCron *cron.Cron
func initCron() error {
	MyCron=cron.New(cron.WithSeconds())//支持秒级定时器
	_,err:=MyCron.AddFunc("0/3 * * * * *", FailTrans)
	return err
}

const FailSql="update translog set STATUS=2 where TIMESTAMPDIFF(SECOND,updatetime,now())>20 and STATUS<>2"

//定时取消交易
func FailTrans()  {
	 _,err:=Trans.GetDB().Exec(FailSql)
	 if err!=nil{
	 	log.Println(err)
	 }
}
func main()  {
	c:=make(chan error)
	go func() {
		err:=Trans.DBInit("a")
		if err!=nil{
			c<-err
		}
	}()
	go func() {
		err:=initCron()
		if err!=nil{
			c<-err
		}
		MyCron.Start()//开启定时任务
	}()
	err:=<-c
	log.Fatal(err)

}

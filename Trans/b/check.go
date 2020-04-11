package main

import (
	"encoding/json"
	"flag"
	"github.com/streadway/amqp"
	"log"
	"go-mq/Lib"
	"go-mq/Trans"
)
const saveSql="insert into translog(tid,`from`,`to`,money,updatetime) values(?,?,?,?,now())"
func saveLog(tm *Trans.TransModel){
	_,err:=Trans.GetDB().Exec(saveSql,tm.Tid,tm.From,tm.To,tm.Money)
	if err!=nil{
		log.Println(err)
	}
}
func recFromA(msgs <-chan amqp.Delivery ,c string ){
	for msg:=range msgs{
		 tm:=Trans.NewTransModel()
		err:=json.Unmarshal(msg.Body,tm)
		if err!=nil{
			log.Println(err)
		}else {
			go func(t *Trans.TransModel) {
				defer msg.Ack(false )
				saveLog(t)
			}(tm)
		}
	}
}
var myclient *Lib.MQ  //暴露myclient
func main()  {
	var c *string
	c=flag.String("c","","消费者名称")
	flag.Parse()
	if *c==""{
		log.Fatal("c参数一定要写")
	}
	dberr:=Trans.DBInit("b") //DB 初始化
	if dberr!=nil{
		log.Fatal("DB error:",dberr)
	}
	myclient=Lib.NewMQ()
	err:=myclient.Channel.Qos(2,0,false)
	if err!=nil{
		log.Fatal(err)
	}
	myclient.Consume(Lib.QUEUE_TRANS,*c,recFromA)

	defer myclient.Channel.Close()
}
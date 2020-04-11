package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"go-mq/Trans"
)

func main()  {
	router:=gin.Default()
	router.Use(Trans.ErrorMiddleware())
	router.Handle("POST","/", func(context *gin.Context) {
		tm:=Trans.NewTransModel()
		err:=context.BindJSON(&tm)
		Trans.CheckError(err,"参数失败:")

		//执行转账--A公司
		err=Trans.TransMoney(tm)
		Trans.CheckError(err,"转账失败-A:")

        context.JSON(200,gin.H{"result":tm.String()})

	})

	c:=make(chan error)
	go func() {
		err:=router.Run(":8088")
		if err!=nil{
			c<-err
		}
	}()
	go func() {
		err:=Trans.DBInit("a")
		if err!=nil{
			c<-err
		}
	}()
	err:=<-c
	log.Fatal(err)
}
package AppInit

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)
var MQConn *amqp.Connection
func init()  {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d/", "shenyi", "123", "127.0.0.1", 5672)
	conn, err := amqp.Dial(dsn)
	if err!=nil{
		log.Fatal(err)
	}
	MQConn=conn

}
func GetConn() *amqp.Connection {
	return MQConn
}
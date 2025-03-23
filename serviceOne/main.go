package main

import (
	"log"
	"net/http"
	rabbitmq "rabbitmq/serviceOne/internal/rabbitMq"
	internal "rabbitmq/serviceOne/internal/service"
)
func main(){
	var rabbit rabbitmq.RabbitMQ
	router := http.NewServeMux()
	conn,err := rabbit.NewConnection("amqp://guest:guest@localhost:5672/")
	if err != nil{
		log.Printf("%s",err)
	}
	defer conn.Close()
	ch,q := rabbit.OpenChannel(conn,"firstQueue")
	_ = internal.NewHandler(router,q,ch)
	_ = http.ListenAndServe("8080",router)
}
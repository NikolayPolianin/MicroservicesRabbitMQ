package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)
func main(){
	conn,err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil{
		log.Fatalf("unable to open connect to RabbitMQ server. Error: %s",err)
	}
	defer conn.Close()

	ch,err := conn.Channel()
	if err != nil{
		log.Fatalf("failed to open channel. Error: %s", err)
	}
	defer ch.Close()
	q,err := ch.QueueDeclare(
		"firstQueue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil{
		log.Fatalf("failed to create queue. Error: %s", err)
	}
	ctx,cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	body := "hello world"
	err = ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType:"text/plain",
			Body: []byte(body)})
	if err != nil{
		log.Fatalf("failed publish message Error: %s",err)
	}
	
}
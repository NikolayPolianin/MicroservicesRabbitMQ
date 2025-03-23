package rabbitmq

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)
type RabbitMQ struct{
}

func (rabbit *RabbitMQ)NewConnection(path string)(*amqp.Connection,error){
	conn,err := amqp.Dial(path)
	if err != nil{
		log.Fatalf("unable to open connect to RabbitMQ server. Error: %s",err)
	}
	defer conn.Close()
	return conn,nil
}
func(rabbit *RabbitMQ)OpenChannel(conn *amqp.Connection,name string)(*amqp.Channel,*amqp.Queue){
	ch,err := conn.Channel()
	if err != nil{
		log.Fatalf("failed to open channel. Error: %s", err)
	}
	defer ch.Close()
	q,err := ch.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil{
		log.Fatalf("failed to create queue. Error: %s", err)
	}
	return ch,&q
}
func(rabbit *RabbitMQ)SendMessageInChannel(ch *amqp.Channel,name string,body string)error{
	ctx,cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	err := ch.PublishWithContext(ctx,
		"",
		name,
		false,
		false,
		amqp.Publishing{
			ContentType:"text/plain",
			Body: []byte(body)})
	if err != nil{
		log.Fatalf("failed publish message Error: %s",err)
	}
	return nil
}
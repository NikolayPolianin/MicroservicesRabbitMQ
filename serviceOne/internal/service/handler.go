package internal

import (
	"encoding/json"
	"log"
	"net/http"
	rabbitmq "rabbitmq/serviceOne/internal/rabbitMq"
	"rabbitmq/serviceOne/internal/service/dto"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Handler struct{
	http.ServeMux
	ch *amqp.Queue
	q *amqp.Channel
}

func NewHandler(router *http.ServeMux,ch *amqp.Queue,q *amqp.Channel)*Handler{
	handler := Handler{
		ServeMux: *router,
		ch:ch,
		q:q,
	}
	router.HandleFunc("POST /home",handler.HelloWorld(q))
	return &handler
}

func (h *Handler)HelloWorld(ch *amqp.Channel)http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		var message dto.Message
		var rabbit rabbitmq.RabbitMQ
		_ = json.NewDecoder(r.Body).Decode(&message)
		err := rabbit.SendMessageInChannel(ch,"firstQueue",message.Message)
		if err != nil{
			log.Printf("%s",err)
		}
	}

}


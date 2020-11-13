package mq

import (
	"encoding/json"
	"fmt"
	"log"

	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
)

// RabbitChan communicates with mq channel
type rabbitChan struct {
	Chan *amqp.Channel
}

// RabbitChan keeps connection to MQ Channel
var RabbitChan DefaultExchange

func InitRabbit() {
	RabbitChan = Chan.rabbitChan
}

func (r *rabbitChan) Test() error {
	return nil
}

func (r *rabbitChan) Listen() error {
	// todo
	return nil
}

func (r *rabbitChan) RPC(routeKey string, data interface{}) (res *interface{}, err error) {
	u1, err := uuid.NewV1()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return nil, err
	}
	corrID := u1.String()

	body, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return nil, err
	}

	// Declare Queue if does not exist
	q, err := r.Chan.QueueDeclare(
		"",    // name
		false, // durable
		true,  // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return nil, err
	}

	msgs, err := r.Chan.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return nil, err
	}

	// Publish request
	err = r.Chan.Publish(
		"",       // exchange
		routeKey, // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: corrID,
			ReplyTo:       q.Name,
			Body:          body,
		})
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return nil, err
	}

	for d := range msgs {
		if corrID == d.CorrelationId {
			err := json.Unmarshal(d.Body, res)
			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}
			break
		}
	}
	return
}

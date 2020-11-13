package mq

import "errors"

type RabbitMQError error

var (
	err = RabbitMQError(errors.New("Error"))
)

func IsRabbitMQError(err error) bool {
	_, ok := err.(RabbitMQError)
	return ok
}

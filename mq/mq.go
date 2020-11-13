package mq

import (
	"fmt"

	config "gitlab.com/logitab/back-end-team/location-tracker-go/config"

	"github.com/streadway/amqp"
)

// ChanConnector keeps connection to all MQs
type ChanConnector struct {
	Chan       *amqp.Channel
	rabbitChan *rabbitChan
}

var (
	// Chan stores mq connection
	Chan *ChanConnector
)

// Init intializes
func Init() {
	channel, err := Connect()
	if err != nil || channel == nil {
		panic("Cannot init channel")
	}
	Chan = channel
	InitRabbit()
}

// Connect connects to a mq.
func Connect() (*ChanConnector, error) {
	cfg := config.Cfg.RabbitMQ

	// Connaect to Rabbit Mq
	conn, err := amqp.Dial(fmt.Sprintf("%s://%s:%s@%s:%s/",
		cfg.Protocol, cfg.Username, cfg.Password, cfg.IP, cfg.Port))
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to RabbitMQ!")

	// Get channel instance
	ch, err := conn.Channel()
	s := &ChanConnector{
		Chan:       ch,
		rabbitChan: &rabbitChan{ch},
	}

	return s, nil
}

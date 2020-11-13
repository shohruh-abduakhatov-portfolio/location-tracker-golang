package mq

// DefaultExchange is an interface to work with mongodb.
type DefaultExchange interface {
	Test() error
	Listen() error
	RPC(routeKey string, data interface{}) (res *interface{}, err error)
}

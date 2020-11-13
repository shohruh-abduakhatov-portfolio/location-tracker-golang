package executor

import "gitlab.com/logitab/back-end-team/location-tracker-go/mq"

type UserClient interface {
	VerifyToken(token string) (*User, error)
}

type UserService struct {
	*mq.ChanConnector
}

type User struct {
	ID string `json:"id"`
}

func NewUserService(conn *mq.ChanConnector) UserClient {
	return &UserService{conn}
}

func (s *UserService) VerifyToken(token string) (*User, error) {
	res, err := mq.RabbitChan.RPC(getCallingFunc(), nil)
	if err != nil {
		return nil, err
	}

	field, ok := (*res).(*User)
	if !ok {
		return nil, errUnexpectedField
	}

	return field, nil
}

package ec

import (
	"net/rpc"

	"github.com/sirupsen/logrus"
)

// Service is a client to the email notification microservice
type Service interface {
	Send(alias string, payload map[string]interface{}) (*SendTransactionlEmailResponse, error)
}

type service struct {
	address string
	logger  *logrus.Logger
}

// New returns a new service
func New(address string, logger *logrus.Logger) Service {
	return &service{address, logger}
}

// Send sends an email to the microservice
func (s *service) Send(alias string, payload map[string]interface{}) (*SendTransactionlEmailResponse, error) {
	client, err := rpc.DialHTTP("tcp", s.address)
	if err != nil {
		return &SendTransactionlEmailResponse{}, err
	}

	var response SendTransactionlEmailResponse
	request := SendTransactionlEmailRequest{Alias: alias, Payload: payload}
	err = client.Call("RPC.SendEmail", request, &response)
	if err != nil {
		return &SendTransactionlEmailResponse{}, err
	}

	return &response, nil
}

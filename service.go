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
	discoverService DiscoverService
	logger          *logrus.Logger
}

// New returns a new service
func New(discoverService DiscoverService, logger *logrus.Logger) Service {
	return &service{discoverService, logger}
}

// Send sends an email to the microservice
func (s *service) Send(alias string, payload map[string]interface{}) (*SendTransactionlEmailResponse, error) {
	address, err := s.discoverService.GetService("/services/email")
	if err != nil {
		return &SendTransactionlEmailResponse{}, err
	}

	client, err := rpc.DialHTTP("tcp", address)
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

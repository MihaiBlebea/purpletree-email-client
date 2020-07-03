package ec

// SpyService is a client to the email notification microservice
type spyService struct {
	requests  []*SendTransactionlEmailRequest
	responses []*SendTransactionlEmailResponse
}

// NewSpyService returns a new spy service for testing
func NewSpyService() Service {
	return &spyService{}
}

// Send sends an email to the microservice
func (s *spyService) Send(alias string, payload map[string]interface{}) (*SendTransactionlEmailResponse, error) {
	s.requests = append(s.requests, &SendTransactionlEmailRequest{
		Alias:   alias,
		Payload: payload,
	})
	response := SendTransactionlEmailResponse{Success: true}
	s.responses = append(s.responses, &response)

	return &response, nil
}

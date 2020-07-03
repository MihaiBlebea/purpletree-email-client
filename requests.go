package ec

// SendTransactionlEmailRequest request contract
type SendTransactionlEmailRequest struct {
	Alias   string
	Payload map[string]interface{}
}

// SendTransactionlEmailResponse response contract
type SendTransactionlEmailResponse struct {
	Success bool
}

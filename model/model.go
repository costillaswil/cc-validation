package model

// InputParameter - http json request parameter model
type InputParameter struct {
	CardNumber string `json:"cardNumber"`
}

// Response - http response model
type Response struct {
	CardNumber   string `json:"cardNumber"`
	CardType     string `json:"cardType"`
	Status       string `json:"status"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

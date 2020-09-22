package erros

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status    int    `json:"status"`
	Error   string `json:"error"`
}


func NewBadRequest(message string) *RestErr{
  return &RestErr{
  	Message: message,
  	Status: http.StatusBadRequest,
  	Error: "bad_request",
  }
}

func NewNotFound(message string) *RestErr{
	return &RestErr{
		Message: message,
		Status: http.StatusNotFound,
		Error: "not found",
	}
}
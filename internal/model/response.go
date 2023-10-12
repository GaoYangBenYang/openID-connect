package model

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(code int, message string, data interface{}) *response {
	return &response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

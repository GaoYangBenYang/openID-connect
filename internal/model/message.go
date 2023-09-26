package model

type Message struct {
	Code    int
	Message string
	Data    interface{}
}

func NewMessage(code int, message string, data interface{}) *Message {
	return &Message{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

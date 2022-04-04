package models

//响应体数据
type Response struct {
	Code string
	Message string
	Data interface{}
}
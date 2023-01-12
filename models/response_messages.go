package models

type ResponseMessage struct {
	Code    int    `json:"code" example:"302"`
	Message string `json:"message" example:"message"`
}

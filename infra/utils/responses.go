package utils

import "github.com/gin-gonic/gin"

//Responses it holds all required properties to return a response.
type Responses struct {
	StatusCode int         `json:"code"`
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

//APIResponse it constructs the HTTP responses returned to the client
func APIResponse(ctx *gin.Context, message string, statusCode int, status bool, data interface{}) {

	jsonResponse := Responses{
		StatusCode: statusCode,
		Status:     status,
		Message:    message,
		Data:       data,
	}

	if statusCode >= 400 {
		ctx.JSON(statusCode, jsonResponse)
		defer ctx.AbortWithStatus(statusCode)
	} else {
		ctx.JSON(statusCode, jsonResponse)
	}
}

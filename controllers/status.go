package controllers

import "github.com/gin-gonic/gin"

type ApiResponse struct {
	Code          int         `json:"code"`
	Status        string      `json:"status"`
	Message       string      `json:"message"`
	MessageDetail string      `json:"message_detail"`
	Data          interface{} `json:"data"`
}

func StatusUpdate(c *gin.Context) {

}

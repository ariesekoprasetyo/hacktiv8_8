package controllers

import (
	"Assigment_8/status"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiResponse struct {
	Code          int         `json:"code"`
	Status        string      `json:"status"`
	Message       string      `json:"message"`
	MessageDetail string      `json:"message_detail"`
	Data          interface{} `json:"data"`
}

func UpdateStatus(c *gin.Context) {
	var statusRequest status.StatusReq
	if err := c.ShouldBindJSON(&statusRequest); err != nil {
		response := ApiResponse{
			Code:          http.StatusBadRequest,
			Status:        "Gagal Bind Json",
			Message:       "Gagal Update Status Water & Wind",
			MessageDetail: err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err := status.StatusUpdate(statusRequest)
	if err != nil {
		response := ApiResponse{
			Code:          http.StatusBadRequest,
			Status:        "Gagal",
			Message:       "Gagal Update Status Water & Wind",
			MessageDetail: err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	c.JSON(http.StatusOK, ApiResponse{
		Code:          http.StatusOK,
		Status:        "Berhasil",
		Message:       "Berhasil Update Status Water & Wind",
		MessageDetail: "",
		Data:          statusRequest,
	})
	fmt.Println(statusRequest)
}

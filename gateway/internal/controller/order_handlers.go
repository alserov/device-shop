package controller

import (
	"github.com/alserov/device-shop/gateway/pkg/models"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	"github.com/gin-gonic/gin"
)

type OrderHandler interface {
	CreateOrder(c *gin.Context)
	UpdateOrder(c *gin.Context)
	CheckOrder(c *gin.Context)
}

func (h *handler) CreateOrder(c *gin.Context) {
	var req models.CreateOrderReq

	if err := c.ShouldBindJSON(&req); err != nil {
		responser.UserError(c.Writer, "failed to decode req body")
		return
	}
}

func (h *handler) UpdateOrder(c *gin.Context) {

}

func (h *handler) CheckOrder(c *gin.Context) {

}

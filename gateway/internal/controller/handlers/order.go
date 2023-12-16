package handlers

import (
	"context"

	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/internal/utils/validation"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	"github.com/alserov/device-shop/proto/gen/order"

	"github.com/gin-gonic/gin"
	"log/slog"
	"time"
)

type OrderHandler interface {
	CreateOrder(c *gin.Context)
	UpdateOrder(c *gin.Context)
	CheckOrder(c *gin.Context)
	CancelOrder(c *gin.Context)
}

type orderHandler struct {
	client order.OrdersClient
	log    *slog.Logger
}

func NewOrderHandler(c order.OrdersClient, log *slog.Logger) OrderHandler {
	return &orderHandler{
		client: c,
		log:    log,
	}
}

const (
	invalidQueryParam = "invalid query param"
)

func (h *orderHandler) CreateOrder(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "ordersHandler.CreateOrder"

	order, err := utils.Decode[order.CreateOrderReq](c.Request, validation.CheckCreateOrder)
	if err != nil {
		w.UserError(err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	res, err := h.client.CreateOrder(ctx, order)
	if err != nil {
		w.HandleServiceError(err, op, h.log)
		return
	}

	w.Value(res)
}

func (h *orderHandler) UpdateOrder(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "ordersHandler.UpdateOrder"

	orderStatus, err := utils.Decode[order.UpdateOrderReq](c.Request, validation.CheckUpdateOrder)
	if err != nil {
		w.UserError(err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	_, err = h.client.UpdateOrder(ctx, orderStatus)
	if err != nil {
		w.HandleServiceError(err, op, h.log)
		return
	}

	w.OK()
}

func (h *orderHandler) CancelOrder(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "ordersHandler.CancelOrder"

	orderUUID := c.Param("order_uuid")
	if orderUUID == "" {
		w.UserError(invalidQueryParam)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	_, err := h.client.CancelOrder(ctx, &order.CancelOrderReq{OrderUUID: orderUUID})
	if err != nil {
		w.HandleServiceError(err, op, h.log)
		return
	}

	w.OK()
}

func (h *orderHandler) CheckOrder(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "ordersHandler.CheckOrder"

	orderUUID := c.Param("order_uuid")
	if orderUUID == "" {
		w.UserError(invalidQueryParam)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	order, err := h.client.CheckOrder(ctx, &order.CheckOrderReq{OrderUUID: orderUUID})
	if err != nil {
		w.HandleServiceError(err, op, h.log)
		return
	}

	w.Data(responser.H{
		"status":    order.Status,
		"price":     order.Price,
		"createdAt": order.CreatedAt.AsTime(),
		"devices":   order.Devices,
	})
}

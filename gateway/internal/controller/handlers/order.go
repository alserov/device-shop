package handlers

import (
	"context"
	"github.com/alserov/device-shop/gateway/internal/logger"

	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/internal/utils/validation"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	"github.com/alserov/device-shop/proto/gen/order"

	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"time"
)

type OrdersHandler interface {
	CreateOrder(c *gin.Context)
	UpdateOrder(c *gin.Context)
	CheckOrder(c *gin.Context)
}

func NewOrderHandler(orderAddr string, log *slog.Logger) OrdersHandler {
	return &ordersHandler{
		serviceAddr: orderAddr,
		log:         log,
	}
}

const (
	invalidQueryParam = "invalid query param"
)

type ordersHandler struct {
	serviceAddr string
	log         *slog.Logger
}

func (h *ordersHandler) CreateOrder(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "ordersHandler.CreateOrder"

	order, err := utils.Decode[order.CreateOrderReq](c.Request, validation.CheckCreateOrder)
	if err != nil {
		w.UserError(err.Error())
		return
	}

	cl, cc, err := client.DialOrder(h.serviceAddr)
	if err != nil {
		h.log.Error("failed to dial order service", logger.Error(err, op))
		w.ServerError()
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	res, err := cl.CreateOrder(ctx, order)
	if err != nil {
		w.HandleServiceError(err, "cl.CreateOrder", h.log)
		return
	}

	w.Value(res)
}

func (h *ordersHandler) UpdateOrder(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "ordersHandler.UpdateOrder"

	orderStatus, err := utils.Decode[order.UpdateOrderReq](c.Request, validation.CheckUpdateOrder)
	if err != nil {
		w.UserError(err.Error())
		return
	}

	cl, cc, err := client.DialOrder(h.serviceAddr)
	if err != nil {
		h.log.Error("failed to dial order service", logger.Error(err, op))
		w.ServerError()
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	_, err = cl.UpdateOrder(ctx, orderStatus)
	if err != nil {
		w.HandleServiceError(err, "cl.UpdateOrder", h.log)
		return
	}

	c.Status(http.StatusOK)
}

func (h *ordersHandler) CheckOrder(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "ordersHandler.CheckOrder"

	orderUUID := c.Param("orderUUID")
	if orderUUID == "" {
		w.UserError(invalidQueryParam)
		return
	}

	cl, cc, err := client.DialOrder(h.serviceAddr)
	if err != nil {
		h.log.Error("failed to dial order service", logger.Error(err, op))
		w.ServerError()
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	order, err := cl.CheckOrder(ctx, &order.CheckOrderReq{OrderUUID: orderUUID})
	if err != nil {
		w.HandleServiceError(err, "cl.CheckOrder", h.log)
		return
	}

	w.Data(responser.H{
		"status":    order.Status,
		"price":     order.Price,
		"createdAt": order.CreatedAt.AsTime(),
		"devices":   order.Devices,
	})
}

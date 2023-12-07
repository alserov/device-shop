package handlers

import (
	"context"
	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/internal/utils/validation"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	"github.com/alserov/device-shop/proto/gen/order"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
	"log/slog"
	"net/http"
	"time"
)

type OrdersHandler interface {
	CreateOrder(c *gin.Context)
	UpdateOrder(c *gin.Context)
	CheckOrder(c *gin.Context)
}

func NewOrderHandler(orderAddr string, logger *slog.Logger) OrdersHandler {
	return &ordersHandler{
		orderAddr: orderAddr,
		logger:    logger,
	}
}

type ordersHandler struct {
	orderAddr string
	logger    *slog.Logger
}

func (h *ordersHandler) CreateOrder(c *gin.Context) {
	order, err := utils.Decode[order.CreateOrderReq](c.Request, validation.CheckCreateOrder)
	if err != nil {
		responser.UserError(c.Writer, err.Error())
		return
	}

	cl, cc, err := client.DialOrder(h.orderAddr)
	if err != nil {
		responser.ServerError(c.Writer, h.logger, err)
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	res, err := cl.CreateOrder(ctx, order)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			responser.UserError(c.Writer, st.Message())
			return
		}
		responser.ServerError(c.Writer, h.logger, err)
		return
	}

	responser.Data(c.Writer, responser.H{
		"orderUUID": res.OrderUUID,
	})
}

func (h *ordersHandler) UpdateOrder(c *gin.Context) {
	orderStatus, err := utils.Decode[order.UpdateOrderReq](c.Request, validation.CheckUpdateOrder)
	if err != nil {
		responser.UserError(c.Writer, err.Error())
		return
	}

	cl, cc, err := client.DialOrder(h.orderAddr)
	if err != nil {
		responser.ServerError(c.Writer, h.logger, err)
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	_, err = cl.UpdateOrder(ctx, orderStatus)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			responser.UserError(c.Writer, st.Message())
			return
		}
		responser.ServerError(c.Writer, h.logger, err)
		return
	}

	c.Status(http.StatusOK)
}

func (h *ordersHandler) CheckOrder(c *gin.Context) {
	orderUUID := c.Param("orderUUID")
	if orderUUID == "" {
		responser.UserError(c.Writer, "invalid orderUUID param")
		return
	}

	cl, cc, err := client.DialOrder(h.orderAddr)
	if err != nil {
		responser.ServerError(c.Writer, h.logger, err)
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	order, err := cl.CheckOrder(ctx, &order.CheckOrderReq{OrderUUID: orderUUID})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			responser.UserError(c.Writer, st.Message())
			return
		}
		responser.ServerError(c.Writer, h.logger, err)
		return
	}

	responser.Data(c.Writer, responser.H{
		"status":    order.Status,
		"price":     order.Price,
		"createdAt": order.CreatedAt.AsTime(),
		"devices":   order.Devices,
	})
}

package controller

import (
	"context"
	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	"github.com/alserov/device-shop/order-service/pkg/entity"
	pb "github.com/alserov/device-shop/proto/gen"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
)

type Orderer interface {
	CreateOrder(c *gin.Context)
	UpdateOrder(c *gin.Context)
	CheckOrder(c *gin.Context)
}

func (h *handler) CreateOrder(c *gin.Context) {
	msg, err := utils.RequestToPBMessage[entity.CreateOrderReq, pb.CreateOrderReq](c.Request, utils.CreateOrderToPB)
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

	res, err := cl.CreateOrder(ctx, msg)
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

func (h *handler) UpdateOrder(c *gin.Context) {
	msg, err := utils.RequestToPBMessage[entity.UpdateOrderReq, pb.UpdateOrderReq](c.Request, utils.UpdateOrderToPB)
	if err != nil {
		responser.ServerError(c.Writer, h.logger, err)
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

	_, err = cl.UpdateOrder(ctx, msg)
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

func (h *handler) CheckOrder(c *gin.Context) {
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

	order, err := cl.CheckOrder(ctx, &pb.CheckOrderReq{OrderUUID: orderUUID})
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

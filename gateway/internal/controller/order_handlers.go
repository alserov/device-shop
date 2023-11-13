package controller

import (
	"context"
	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/gateway/pkg/models"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	pb "github.com/alserov/device-shop/proto/gen"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
	"net/http"
	"os"
	"time"
)

type OrderHandler interface {
	CreateOrder(c *gin.Context)
	UpdateOrder(c *gin.Context)
	CheckOrder(c *gin.Context)
}

var (
	ORDER_ADDR = os.Getenv("ORDER_ADDR")
)

func (h *handler) CreateOrder(c *gin.Context) {
	msg, err := utils.RequestToPBMessage[models.CreateOrderReq, pb.CreateOrderReq](c.Request, utils.CreateOrderToPB)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}

	cl, cc, err := client.DialOrder(ORDER_ADDR)
	if err != nil {
		responser.ServerError(c.Writer, err)
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
		responser.ServerError(c.Writer, err)
		return
	}

	responser.Data(c.Writer, responser.H{
		"orderUUID": res.OrderUUID,
	})
}

func (h *handler) UpdateOrder(c *gin.Context) {
	msg, err := utils.RequestToPBMessage[models.UpdateOrderReq, pb.UpdateOrderReq](c.Request, utils.UpdateOrderToPB)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}

	cl, cc, err := client.DialOrder(ORDER_ADDR)
	if err != nil {
		responser.ServerError(c.Writer, err)
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
		responser.ServerError(c.Writer, err)
		return
	}

	c.Status(http.StatusOK)
}

func (h *handler) CheckOrder(c *gin.Context) {
	msg, err := utils.RequestToPBMessage[models.CheckOrderReq, pb.CheckOrderReq](c.Request, utils.CheckOrderToPB)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}

	cl, cc, err := client.DialOrder(ORDER_ADDR)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	_, err = cl.CheckOrder(ctx, msg)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			responser.UserError(c.Writer, st.Message())
			return
		}
		responser.ServerError(c.Writer, err)
		return
	}

	c.Status(http.StatusOK)
}

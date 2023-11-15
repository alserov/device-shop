package controller

import (
	"context"
	"github.com/alserov/device-shop/device-service/pkg/entity"
	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/gateway/pkg/models"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	"github.com/alserov/device-shop/proto/gen"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
)

type AdminHandler interface {
	CreateDevice(c *gin.Context)
	DeleteDevice(c *gin.Context)
	UpdateDevice(c *gin.Context)
}

func (h *handler) CreateDevice(c *gin.Context) {
	msg, err := utils.RequestToPBMessage[entity.Device, pb.CreateReq](c.Request, utils.CreateDeviceToPB)
	if err != nil {
		responser.ServerError(c.Writer, h.logger, err)
		return
	}

	cl, cc, err := client.DialDevice(DEVICE_ADDR)
	if err != nil {
		responser.ServerError(c.Writer, h.logger, err)
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(1000)*time.Millisecond)
	defer cancel()

	_, err = cl.CreateDevice(ctx, msg)
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

func (h *handler) DeleteDevice(c *gin.Context) {
	deviceUUID := c.Param("deviceUUID")

	if deviceUUID == "" {
		responser.UserError(c.Writer, "invalid device uuid value")
		return
	}

	cl, cc, err := client.DialDevice(DEVICE_ADDR)
	if err != nil {
		responser.ServerError(c.Writer, h.logger, err)
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(1000)*time.Millisecond)
	defer cancel()

	_, err = cl.DeleteDevice(ctx, &pb.DeleteReq{UUID: deviceUUID})
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

func (h *handler) UpdateDevice(c *gin.Context) {
	msg, err := utils.RequestToPBMessage[models.UpdateDeviceReq, pb.UpdateReq](c.Request, utils.UpdateDeviceToPB)
	if err != nil {
		responser.ServerError(c.Writer, h.logger, err)
		return
	}

	cl, cc, err := client.DialDevice(DEVICE_ADDR)
	if err != nil {
		responser.ServerError(c.Writer, h.logger, err)
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(1000)*time.Millisecond)
	defer cancel()

	_, err = cl.UpdateDevice(ctx, msg)
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

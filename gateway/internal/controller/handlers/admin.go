package handlers

import (
	"context"
	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/internal/utils/validation"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	pb "github.com/alserov/device-shop/proto/gen"
	"github.com/alserov/device-shop/proto/gen/admin"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
	"log/slog"
	"net/http"
	"time"
)

type AdminHandler interface {
	CreateDevice(c *gin.Context)
	DeleteDevice(c *gin.Context)
	UpdateDevice(c *gin.Context)
}

func NewAdminHandler(deviceAddr, userAddr string, logger *slog.Logger) AdminHandler {
	return &adminHandler{
		logger:     logger,
		deviceAddr: deviceAddr,
		userAddr:   userAddr,
	}
}

type adminHandler struct {
	deviceAddr string
	userAddr   string
	logger     *slog.Logger
}

func (h *adminHandler) CreateDevice(c *gin.Context) {
	device, err := utils.Decode[admin.CreateDeviceReq](c.Request, validation.CheckCreateDevice)
	if err != nil {
		responser.UserError(c.Writer, err.Error())
		return
	}

	cl, cc, err := client.DialDevice(h.deviceAddr)
	if err != nil {
		responser.ServerError(c.Writer, h.logger, err)
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(1000)*time.Millisecond)
	defer cancel()

	_, err = cl.CreateDevice(ctx, device)
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

func (h *adminHandler) DeleteDevice(c *gin.Context) {
	deviceUUID := c.Param("deviceUUID")

	if deviceUUID == "" {
		responser.UserError(c.Writer, "invalid device uuid value")
		return
	}

	cl, cc, err := client.DialDevice(h.deviceAddr)
	if err != nil {
		responser.ServerError(c.Writer, h.logger, err)
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(1000)*time.Millisecond)
	defer cancel()

	_, err = cl.DeleteDevice(ctx, &pb.DeleteDeviceReq{UUID: deviceUUID})
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

func (h *adminHandler) UpdateDevice(c *gin.Context) {
	device, err := utils.Decode[pb.UpdateDeviceReq](c.Request, validation.CheckUpdateDevice)
	if err != nil {
		responser.UserError(c.Writer, err.Error())
		return
	}

	cl, cc, err := client.DialDevice(h.userAddr)
	if err != nil {
		responser.ServerError(c.Writer, h.logger, err)
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(1000)*time.Millisecond)
	defer cancel()

	_, err = cl.UpdateDevice(ctx, device)
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

package handlers

import (
	"context"
	"github.com/alserov/device-shop/gateway/internal/logger"
	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/internal/utils/validation"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	"github.com/alserov/device-shop/proto/gen/device"

	"github.com/gin-gonic/gin"
	"log/slog"
	"time"
)

func handleServiceError(msg string) {

}

type AdminHandler interface {
	CreateDevice(c *gin.Context)
	DeleteDevice(c *gin.Context)
	UpdateDevice(c *gin.Context)
	UpdateDeviceAmount(c *gin.Context)
}

type AdminH struct {
	DeviceAddr string
	Log        *slog.Logger
}

func NewAdminHandler(ah *AdminH) AdminHandler {
	return &adminHandler{
		log:         ah.Log,
		serviceAddr: ah.DeviceAddr,
	}
}

type adminHandler struct {
	serviceAddr string
	log         *slog.Logger
}

func (h *adminHandler) UpdateDeviceAmount(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "increase device amount"

	deviceUUIDandAmount, err := utils.Decode[device.IncreaseDeviceAmountByUUIDReq](c.Request)
	if err != nil {
		w.UserError(err.Error())
		return
	}

	cl, cc, err := client.DialDevice(h.serviceAddr)
	if err != nil {
		h.log.Error("failed to dial device service", logger.Error(err, op))
		w.ServerError()
		return
	}
	defer cc.Close()

	_, err = cl.IncreaseDeviceAmount(c.Request.Context(), deviceUUIDandAmount)
	if err != nil {
		w.HandleServiceError(err, "cl.IncreaseDeviceAmount", h.log)
		return
	}

	w.OK()
}

func (h *adminHandler) CreateDevice(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "adminHandler.CreateDevice"

	device, err := utils.Decode[device.CreateDeviceReq](c.Request, validation.CheckCreateDevice)
	if err != nil {
		w.UserError(err.Error())
		return
	}

	cl, cc, err := client.DialDevice(h.serviceAddr)
	if err != nil {
		h.log.Error("failed to dial device service", logger.Error(err, op))
		w.ServerError()
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(1000)*time.Millisecond)
	defer cancel()

	_, err = cl.CreateDevice(ctx, device)
	if err != nil {
		w.HandleServiceError(err, "cl.CreateDevice", h.log)
		return
	}

	w.OK()
}

func (h *adminHandler) DeleteDevice(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "adminHandler.DeleteDevice"

	deviceUUID := c.Param("deviceUUID")
	if deviceUUID == "" {
		w.UserError("device uuid is empty")
		return
	}

	cl, cc, err := client.DialDevice(h.serviceAddr)
	if err != nil {
		h.log.Error("failed to dial device service", logger.Error(err, op))
		w.ServerError()
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(1000)*time.Millisecond)
	defer cancel()

	_, err = cl.DeleteDevice(ctx, &device.DeleteDeviceReq{UUID: deviceUUID})
	if err != nil {
		w.HandleServiceError(err, "cl.DeleteDevice", h.log)
		return
	}

	w.OK()
}

func (h *adminHandler) UpdateDevice(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "adminHandler.UpdateDevice"

	device, err := utils.Decode[device.UpdateDeviceReq](c.Request, validation.CheckUpdateDevice)
	if err != nil {
		w.UserError(err.Error())
		return
	}

	cl, cc, err := client.DialDevice(h.serviceAddr)
	if err != nil {
		h.log.Error("failed to dial device service", logger.Error(err, op))
		w.ServerError()
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(1000)*time.Millisecond)
	defer cancel()

	_, err = cl.UpdateDevice(ctx, device)
	if err != nil {
		w.HandleServiceError(err, "cl.UpdateDevice", h.log)
		return
	}

	w.OK()
}

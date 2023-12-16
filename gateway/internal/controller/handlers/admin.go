package handlers

import (
	"context"

	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/internal/utils/validation"
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

func NewAdminHandler(c device.DevicesClient, log *slog.Logger) AdminHandler {
	return &adminHandler{
		log:    log,
		client: c,
	}
}

type adminHandler struct {
	client device.DevicesClient
	log    *slog.Logger
}

func (h *adminHandler) UpdateDeviceAmount(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "adminHandler.UpdateDeviceAmount"

	deviceUUIDAndAmount, err := utils.Decode[device.IncreaseDeviceAmountReq](c.Request)
	if err != nil {
		w.UserError(err.Error())
		return
	}

	_, err = h.client.IncreaseDeviceAmount(c.Request.Context(), deviceUUIDAndAmount)
	if err != nil {
		w.HandleServiceError(err, op, h.log)
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

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(1000)*time.Millisecond)
	defer cancel()

	_, err = h.client.CreateDevice(ctx, device)
	if err != nil {
		w.HandleServiceError(err, op, h.log)
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

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(1000)*time.Millisecond)
	defer cancel()

	_, err := h.client.DeleteDevice(ctx, &device.DeleteDeviceReq{UUID: deviceUUID})
	if err != nil {
		w.HandleServiceError(err, op, h.log)
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

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(1000)*time.Millisecond)
	defer cancel()

	_, err = h.client.UpdateDevice(ctx, device)
	if err != nil {
		w.HandleServiceError(err, op, h.log)
		return
	}

	w.OK()
}

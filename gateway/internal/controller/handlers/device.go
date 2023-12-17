package handlers

import (
	"context"
	"fmt"
	"github.com/alserov/device-shop/gateway/internal/broker"
	"github.com/alserov/device-shop/gateway/internal/cache"
	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/internal/utils/validation"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	"github.com/alserov/device-shop/proto/gen/device"
	"github.com/gin-gonic/gin"
	"log/slog"
	"strconv"
	"strings"
	"time"
)

type DeviceHandler interface {
	GetAllDevices(*gin.Context)
	GetDevicesByTitle(*gin.Context)
	GetDevicesByManufacturer(*gin.Context)
	GetDevicesByPrice(*gin.Context)
	GetDeviceByUUID(*gin.Context)
}

type devicesHandler struct {
	log    *slog.Logger
	client device.DevicesClient
	cache  cache.Repository
	p      broker.MetricsProducer
}

func NewDeviceHandler(c device.DevicesClient, cache cache.Repository, p broker.MetricsProducer, log *slog.Logger) DeviceHandler {
	return &devicesHandler{
		client: c,
		cache:  cache,
		log:    log,
		p:      p,
	}
}

func (h *devicesHandler) GetDeviceByUUID(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "devicesHandler.GetDeviceByUUID"

	uuid, err := utils.Decode[device.GetDeviceByUUIDReq](c.Request)
	if err != nil {
		w.UserError(err.Error())
		return
	}

	device, err := h.client.GetDeviceByUUID(c.Request.Context(), uuid)
	if err != nil {
		w.HandleServiceError(err, op, h.log)
		return
	}

	w.Value(device)
}

func (h *devicesHandler) GetAllDevices(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "devicesHandler.GetAllDevices"

	start := time.Now()

	getDevicesCred, err := utils.Decode[device.GetAllDevicesReq](c.Request, validation.CheckGetAll)
	if err != nil {
		w.UserError(err.Error())
		return
	}

	val, err := h.cache.GetValue(c.Request.Context(), fmt.Sprintf("%d%d", getDevicesCred.Index, getDevicesCred.Amount))
	if err == nil {
		devices, ok := val.([]interface{})
		if ok {
			w.Data(responser.H{
				"data":   val,
				"amount": len(devices),
				"index":  getDevicesCred.Index + 1,
			})
		} else {
			w.Data(responser.H{
				"data":   val,
				"amount": 0,
				"index":  getDevicesCred.Index + 1,
			})
		}
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(1000)*time.Millisecond)
	defer cancel()

	devices, err := h.client.GetAllDevices(ctx, getDevicesCred)
	if err != nil {
		w.HandleServiceError(err, op, h.log)
		return
	}

	err = h.cache.SetValue(c.Request.Context(), &cache.Set{
		Val: devices.Devices,
		Key: fmt.Sprintf("%d%d", getDevicesCred.Index, getDevicesCred.Amount),
	})
	if err != nil {
		h.log.Error("failed to set cache", slog.String("error", err.Error()), slog.String("op", op))
	}

	if err := h.p.Latency(time.Since(start)); err != nil {
		h.log.Error("failed to send message to topic", slog.String("error", err.Error()), slog.String("op", op))
	}

	w.Data(responser.H{
		"data":   devices.Devices,
		"amount": len(devices.Devices),
		"index":  getDevicesCred.Index + 1,
	})
}

func (h *devicesHandler) GetDevicesByTitle(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "devicesHandler.GetDevicesByTitle"

	title := strings.ToLower(c.Param("title"))
	if title == "" {
		w.UserError(invalidQueryParam)
		return
	}

	val, err := h.cache.GetValue(c.Request.Context(), title)
	if err == nil {
		devices, ok := val.([]interface{})
		if ok {
			w.Data(responser.H{
				"data":   val,
				"amount": len(devices),
			})
		} else {
			w.Data(responser.H{
				"data":   val,
				"amount": 0,
			})
		}
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(1000)*time.Millisecond)
	defer cancel()

	devices, err := h.client.GetDevicesByTitle(ctx, &device.GetDeviceByTitleReq{Title: title})
	if err != nil {
		w.HandleServiceError(err, op, h.log)
		return
	}

	err = h.cache.SetValue(c.Request.Context(), &cache.Set{
		Val: devices.Devices,
		Key: title,
	})
	if err != nil {
		h.log.Error("failed to set cache", slog.String("error", err.Error()), slog.String("op", op))
	}

	w.Data(responser.H{
		"data":   devices.Devices,
		"amount": len(devices.Devices),
	})
}

func (h *devicesHandler) GetDevicesByManufacturer(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "devicesHandler.GetDevicesByManufacturer"

	manu := strings.ToLower(c.Param("manu"))
	if manu == "" {
		w.UserError(invalidQueryParam)
		return
	}

	val, err := h.cache.GetValue(c.Request.Context(), manu)
	if err == nil {
		w.Data(responser.H{
			"data":   val,
			"amount": len(val.([]interface{})),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	d, err := h.client.GetDevicesByManufacturer(ctx, &device.GetByManufacturer{
		Manufacturer: manu,
	})
	if err != nil {
		w.HandleServiceError(err, op, h.log)
		return
	}

	err = h.cache.SetValue(c.Request.Context(), &cache.Set{
		Key: manu,
		Val: d.Devices,
	})
	if err != nil {
		h.log.Error("failed to set cache", slog.String("error", err.Error()), slog.String("op", op))
	}

	w.Data(responser.H{
		"data":   d.Devices,
		"amount": len(d.Devices),
	})
}

func (h *devicesHandler) GetDevicesByPrice(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "devicesHandler.GetDevicesByPrice"

	minVal, err := strconv.Atoi(c.Query("min"))
	if err != nil {
		w.UserError("invalid value for 'min' param")
		return
	}

	maxVal, err := strconv.Atoi(c.Query("max"))
	if err != nil {
		w.UserError("invalid value for 'max' param")
		return
	}

	if minVal >= maxVal {
		w.UserError("'min' value can't be equal or greater than 'max' value")
		return
	}

	val, err := h.cache.GetValue(c.Request.Context(), fmt.Sprintf("%d%d", minVal, maxVal))
	if err == nil {
		w.Data(responser.H{
			"data":   val,
			"amount": len(val.([]interface{})),
		})
		return
	}

	r := &device.GetByPrice{
		Min: float32(minVal),
		Max: float32(maxVal),
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	d, err := h.client.GetDevicesByPrice(ctx, r)
	if err != nil {
		w.HandleServiceError(err, op, h.log)
		return
	}

	err = h.cache.SetValue(c.Request.Context(), &cache.Set{
		Key: fmt.Sprintf("%d%d", minVal, maxVal),
		Val: d.Devices,
	})
	if err != nil {
		h.log.Error("failed to set cache", slog.String("error", err.Error()), slog.String("op", op))
	}

	w.Data(responser.H{
		"data":   d.Devices,
		"amount": len(d.Devices),
	})
}

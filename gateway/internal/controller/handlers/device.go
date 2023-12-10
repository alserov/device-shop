package handlers

import (
	"context"
	"fmt"
	"github.com/alserov/device-shop/gateway/internal/logger"
	"github.com/go-redis/redis"

	"github.com/alserov/device-shop/gateway/internal/cache"
	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/internal/utils/validation"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	"github.com/alserov/device-shop/proto/gen/device"
	"github.com/gin-gonic/gin"

	"log/slog"
	"strconv"
	"strings"
	"time"
)

type DevicesHandler interface {
	GetAllDevices(*gin.Context)
	GetDevicesByTitle(*gin.Context)
	GetDevicesByManufacturer(*gin.Context)
	GetDevicesByPrice(*gin.Context)
	GetDeviceByUUID(*gin.Context)
}

type devicesHandler struct {
	serviceAddr string
	cache       cache.Repository
	log         *slog.Logger
}

func NewDevicesHandler(deviceAddr string, rd *redis.Client, log *slog.Logger) DevicesHandler {
	return &devicesHandler{
		serviceAddr: deviceAddr,
		cache:       cache.NewRepo(rd),
		log:         log,
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

	cl, cc, err := client.DialDevice(h.serviceAddr)
	if err != nil {
		h.log.Error("failed to dial device service", logger.Error(err, op))
		w.ServerError()
		return
	}
	defer cc.Close()

	device, err := cl.GetDeviceByUUID(c.Request.Context(), uuid)
	if err != nil {
		w.HandleServiceError(err, "cl.GetDeviceByUUID", h.log)
		return
	}

	w.Value(device)
}

func (h *devicesHandler) GetAllDevices(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "devicesHandler.GetAllDevices"

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

	cl, cc, err := client.DialDevice(h.serviceAddr)
	if err != nil {
		h.log.Error("failed to dial device service", logger.Error(err, op))
		w.ServerError()
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(1000)*time.Millisecond)
	defer cancel()

	devices, err := cl.GetAllDevices(ctx, getDevicesCred)
	if err != nil {
		w.HandleServiceError(err, "cl.GetAllDevices", h.log)
		return
	}

	err = h.cache.SetValue(c.Request.Context(), &cache.Set{
		Val: devices.Devices,
		Key: fmt.Sprintf("%d%d", getDevicesCred.Index, getDevicesCred.Amount),
	})
	if err != nil {
		h.log.Error("failed to set cache", logger.Error(err, "h.cache.SetValue"))
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

	cl, cc, err := client.DialDevice(h.serviceAddr)
	if err != nil {
		h.log.Error("failed to dial device service", logger.Error(err, op))
		w.ServerError()
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(1000)*time.Millisecond)
	defer cancel()

	devices, err := cl.GetDevicesByTitle(ctx, &device.GetDeviceByTitleReq{Title: title})
	if err != nil {
		w.HandleServiceError(err, "cl.GetDevicesByTitle", h.log)
		return
	}

	err = h.cache.SetValue(c.Request.Context(), &cache.Set{
		Val: devices.Devices,
		Key: title,
	})
	if err != nil {
		h.log.Error("failed to set cache", logger.Error(err, "h.cache.SetValue"))
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

	cl, cc, err := client.DialDevice(h.serviceAddr)
	if err != nil {
		h.log.Error("failed to dial device service", logger.Error(err, op))
		w.ServerError()
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	d, err := cl.GetDevicesByManufacturer(ctx, &device.GetByManufacturer{
		Manufacturer: manu,
	})
	if err != nil {
		w.HandleServiceError(err, "cl.GetDevicesByManufacturer", h.log)
		return
	}

	err = h.cache.SetValue(c.Request.Context(), &cache.Set{
		Key: manu,
		Val: d.Devices,
	})
	if err != nil {
		h.log.Error("failed to set cache", logger.Error(err, "h.cache.SetValue"))
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

	cl, cc, err := client.DialDevice(h.serviceAddr)
	if err != nil {
		h.log.Error("failed to dial device service", logger.Error(err, op))
		w.ServerError()
		return
	}
	defer cc.Close()

	r := &device.GetByPrice{
		Min: float32(minVal),
		Max: float32(maxVal),
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	d, err := cl.GetDevicesByPrice(ctx, r)
	if err != nil {
		w.HandleServiceError(err, "cl.GetDevicesByPrice", h.log)
		return
	}

	err = h.cache.SetValue(c.Request.Context(), &cache.Set{
		Key: fmt.Sprintf("%d%d", minVal, maxVal),
		Val: d.Devices,
	})
	if err != nil {
		h.log.Error("failed to set cache", logger.Error(err, "h.cache.SetValue"))
	}

	w.Data(responser.H{
		"data":   d.Devices,
		"amount": len(d.Devices),
	})
}

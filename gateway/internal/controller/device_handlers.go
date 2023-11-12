package controller

import (
	"context"
	"fmt"
	"github.com/alserov/shop/gateway/internal/cache"
	"github.com/alserov/shop/gateway/pkg/client"
	"github.com/alserov/shop/gateway/pkg/models"
	"github.com/alserov/shop/gateway/pkg/responser"
	"github.com/alserov/shop/proto/gen"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type DeviceHandler interface {
	GetAllDevices(c *gin.Context)
	GetDevicesByTitle(c *gin.Context)
	GetDevicesByManufacturer(c *gin.Context)
	GetDevicesByPrice(c *gin.Context)
}

var (
	DEVICE_ADDR = os.Getenv("DEVICE_ADDR")
)

func (h *handler) GetAllDevices(c *gin.Context) {
	var req models.GetAllReq

	val, err := h.cache.GetValue(c.Request.Context(), fmt.Sprintf("%d%d", req.Index, req.Amount))
	if err == nil {
		responser.Data(c.Writer, responser.H{
			"data":   val,
			"amount": len(val.([]interface{})),
			"index":  *req.Index + 1,
		})
		return
	}

	if err = c.ShouldBindJSON(&req); err != nil {
		responser.UserError(c.Writer, "failed to decode req body")
		return
	}

	if err = models.Validate(&req); err != nil {
		responser.UserError(c.Writer, err.Error())
		return
	}

	cl, cc, err := client.DialDevices(DEVICE_ADDR)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}
	defer cc.Close()

	r := &pb.GetAllReq{
		Index:  *req.Index,
		Amount: *req.Amount,
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(1000)*time.Millisecond)
	defer cancel()

	devices, err := cl.GetAllDevices(ctx, r)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			responser.UserError(c.Writer, st.Message())
			return
		}
		responser.ServerError(c.Writer, err)
		return
	}

	err = h.cache.SetValue(c.Request.Context(), &cache.Set{
		Val: devices.Devices,
		Key: fmt.Sprintf("%d%d", req.Index, req.Amount),
	})
	if err != nil {
		log.Println("failed to cache: ", err)
	}

	responser.Data(c.Writer, responser.H{
		"data":   devices.Devices,
		"amount": len(devices.Devices),
		"index":  *req.Index + 1,
	})
}

func (h *handler) GetDevicesByTitle(c *gin.Context) {
	title := strings.ToLower(c.Param("title"))

	val, err := h.cache.GetValue(c.Request.Context(), title)
	if err == nil {
		responser.Data(c.Writer, responser.H{
			"data":   val,
			"amount": len(val.([]interface{})),
		})
		return
	}

	if title == "" {
		responser.UserError(c.Writer, "search value can't be empty string")
		return
	}

	cl, cc, err := client.DialDevices(DEVICE_ADDR)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}
	defer cc.Close()

	r := &pb.GetByTitleReq{
		Title: title,
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(1000)*time.Millisecond)
	defer cancel()

	devices, err := cl.GetDevicesByTitle(ctx, r)

	err = h.cache.SetValue(c.Request.Context(), &cache.Set{
		Val: devices.Devices,
		Key: title,
	})
	if err != nil {
		log.Println("failed to cache: ", err)
	}

	if err != nil {
		if st, ok := status.FromError(err); ok {
			responser.UserError(c.Writer, st.Message())
			return
		}
		responser.ServerError(c.Writer, err)
		return
	}

	responser.Data(c.Writer, responser.H{
		"data":   devices.Devices,
		"amount": len(devices.Devices),
	})
}

func (h *handler) GetDevicesByManufacturer(c *gin.Context) {
	manu := strings.ToLower(c.Param("manu"))
	if manu == "" {
		responser.UserError(c.Writer, "invalid manufacturer")
		return
	}

	val, err := h.cache.GetValue(c.Request.Context(), manu)
	if err == nil {
		responser.Data(c.Writer, responser.H{
			"data":   val,
			"amount": len(val.([]interface{})),
		})
		return
	}

	cl, cc, err := client.DialDevices(DEVICE_ADDR)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}
	defer cc.Close()

	r := &pb.GetByManufacturer{
		Manufacturer: manu,
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	d, err := cl.GetDevicesByManufacturer(ctx, r)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			responser.UserError(c.Writer, st.Message())
			return
		}
		responser.ServerError(c.Writer, err)
		return
	}

	err = h.cache.SetValue(c.Request.Context(), &cache.Set{
		Key: manu,
		Val: d.Devices,
	})
	if err != nil {
		log.Println("failed to cache: ", err)
	}

	responser.Data(c.Writer, responser.H{
		"data":   d.Devices,
		"amount": len(d.Devices),
	})
}

func (h *handler) GetDevicesByPrice(c *gin.Context) {
	minVal, err := strconv.Atoi(c.Query("min"))
	if err != nil {
		responser.UserError(c.Writer, "invalid value for 'min' param")
		return
	}

	maxVal, err := strconv.Atoi(c.Query("max"))
	if err != nil {
		responser.UserError(c.Writer, "invalid value for 'max' param")
		return
	}

	if minVal >= maxVal {
		responser.UserError(c.Writer, "'min' value can't be equal or greater than 'max' value")
		return
	}

	val, err := h.cache.GetValue(c.Request.Context(), fmt.Sprintf("%d%d", minVal, maxVal))
	if err == nil {
		responser.Data(c.Writer, responser.H{
			"data":   val,
			"amount": len(val.([]interface{})),
		})
		return
	}

	cl, cc, err := client.DialDevices(DEVICE_ADDR)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}
	defer cc.Close()

	r := &pb.GetByPrice{
		Min: uint32(minVal),
		Max: uint32(maxVal),
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	d, err := cl.GetDevicesByPrice(ctx, r)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			responser.UserError(c.Writer, st.Message())
			return
		}
		responser.ServerError(c.Writer, err)
		return
	}

	err = h.cache.SetValue(c.Request.Context(), &cache.Set{
		Key: fmt.Sprintf("%d%d", minVal, maxVal),
		Val: d.Devices,
	})
	if err != nil {
		log.Println("failed to cache: ", err)
	}

	responser.Data(c.Writer, responser.H{
		"data":   d.Devices,
		"amount": len(d.Devices),
	})
}

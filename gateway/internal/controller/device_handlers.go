package controller

import (
	"context"
	"fmt"
	"github.com/alserov/device-shop/gateway/internal/cache"
	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	"github.com/alserov/device-shop/proto/gen"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
	"log"
	"strconv"
	"strings"
	"time"
)

type Devicer interface {
	GetAllDevices(*gin.Context)
	GetDevicesByTitle(*gin.Context)
	GetDevicesByManufacturer(*gin.Context)
	GetDevicesByPrice(*gin.Context)
}

func (h *handler) GetAllDevices(c *gin.Context) {
	getDevicesCred, err := utils.Decode[pb.GetAllDevicesReq](c.Request)
	if err != nil {
		responser.UserError(c.Writer, err.Error())
		return
	}

	val, err := h.cache.GetValue(c.Request.Context(), fmt.Sprintf("%d%d", getDevicesCred.Index, getDevicesCred.Amount))
	if err == nil {
		devices, ok := val.([]interface{})
		if ok {
			responser.Data(c.Writer, responser.H{
				"data":   val,
				"amount": len(devices),
				"index":  getDevicesCred.Index + 1,
			})
		} else {
			responser.Data(c.Writer, responser.H{
				"data":   val,
				"amount": 0,
				"index":  getDevicesCred.Index + 1,
			})
		}
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

	devices, err := cl.GetAllDevices(ctx, getDevicesCred)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			responser.UserError(c.Writer, st.Message())
			return
		}
		responser.ServerError(c.Writer, h.logger, err)
		return
	}

	err = h.cache.SetValue(c.Request.Context(), &cache.Set{
		Val: devices.Devices,
		Key: fmt.Sprintf("%d%d", getDevicesCred.Index, getDevicesCred.Amount),
	})
	if err != nil {
		log.Println("failed to cache: ", err)
	}

	responser.Data(c.Writer, responser.H{
		"data":   devices.Devices,
		"amount": len(devices.Devices),
		"index":  getDevicesCred.Index + 1,
	})
}

func (h *handler) GetDevicesByTitle(c *gin.Context) {
	title := strings.ToLower(c.Param("title"))

	val, err := h.cache.GetValue(c.Request.Context(), title)
	if err == nil {
		devices, ok := val.([]interface{})
		if ok {
			responser.Data(c.Writer, responser.H{
				"data":   val,
				"amount": len(devices),
			})
		} else {
			responser.Data(c.Writer, responser.H{
				"data":   val,
				"amount": 0,
			})
		}
		return
	}

	if title == "" {
		responser.UserError(c.Writer, "search value can't be empty string")
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

	devices, err := cl.GetDevicesByTitle(ctx, &pb.GetDeviceByTitleReq{Title: title})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			responser.UserError(c.Writer, st.Message())
			return
		}
		responser.ServerError(c.Writer, h.logger, err)
		return
	}

	err = h.cache.SetValue(c.Request.Context(), &cache.Set{
		Val: devices.Devices,
		Key: title,
	})
	if err != nil {
		log.Println("failed to cache: ", err)
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

	cl, cc, err := client.DialDevice(h.deviceAddr)
	if err != nil {
		responser.ServerError(c.Writer, h.logger, err)
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
		responser.ServerError(c.Writer, h.logger, err)
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

	cl, cc, err := client.DialDevice(h.deviceAddr)
	if err != nil {
		responser.ServerError(c.Writer, h.logger, err)
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
		responser.ServerError(c.Writer, h.logger, err)
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

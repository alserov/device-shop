package controller

import (
	"context"
	"github.com/alserov/shop/api/pkg/client"
	"github.com/alserov/shop/api/pkg/responser"
	"github.com/alserov/shop/device-service/pkg/pb"
	"github.com/alserov/shop/gateway/pkg/models"
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
	var req models.Device

	if err := c.ShouldBindJSON(&req); err != nil {
		responser.UserError(c.Writer, "failed decode req body")
		return
	}

	if err := models.Validate(&req); err != nil {
		responser.UserError(c.Writer, err.Error())
		return
	}

	cl, cc, err := client.DialDevices(DEVICE_ADDR)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}
	defer cc.Close()

	r := &pb.CreateReq{
		Title:        req.Title,
		Description:  req.Description,
		Manufacturer: req.Manufacturer,
		Price:        req.Price,
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(1000)*time.Millisecond)
	defer cancel()

	_, err = cl.CreateDevice(ctx, r)
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

func (h *handler) DeleteDevice(c *gin.Context) {
	deviceUUID := c.Param("deviceUUID")

	if deviceUUID == "" {
		responser.UserError(c.Writer, "invalid device uuid value")
		return
	}

	cl, cc, err := client.DialDevices(DEVICE_ADDR)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}
	defer cc.Close()

	r := &pb.DeleteReq{
		UUID: deviceUUID,
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(1000)*time.Millisecond)
	defer cancel()

	_, err = cl.DeleteDevice(ctx, r)
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

func (h *handler) UpdateDevice(c *gin.Context) {
	var req models.UpdateReq

	if err := c.ShouldBindJSON(&req); err != nil {
		responser.UserError(c.Writer, "failed to decode req body")
		return
	}

	if err := models.Validate(&req); err != nil {
		responser.UserError(c.Writer, err.Error())
		return
	}

	cl, cc, err := client.DialDevices(DEVICE_ADDR)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}
	defer cc.Close()

	r := &pb.UpdateReq{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		UUID:        req.UUID,
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(1000)*time.Millisecond)
	defer cancel()

	_, err = cl.UpdateDevice(ctx, r)
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

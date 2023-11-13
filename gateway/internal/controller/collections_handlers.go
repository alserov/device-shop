package controller

import (
	"context"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/gateway/pkg/models"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	pb "github.com/alserov/shop/proto/gen"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
	"net/http"
	"os"
	"time"
)

type CollectionsHandler interface {
	AddToFavourite(c *gin.Context)
	RemoveFromFavourite(c *gin.Context)
	GetFavourite(c *gin.Context)

	AddToCart(c *gin.Context)
	RemoveFromCart(c *gin.Context)
	GetCart(c *gin.Context)
}

var (
	USER_ADDR = os.Getenv("USER_ADDR")
)

func (h *handler) AddToFavourite(c *gin.Context) {
	var req models.AddReq

	if err := c.ShouldBindJSON(&req); err != nil {
		responser.UserError(c.Writer, "failed to decode req body")
		return
	}

	if err := models.Validate(&req); err != nil {
		responser.UserError(c.Writer, err.Error())
		return
	}

	cl, cc, err := client.DialUser(USER_ADDR)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}
	defer cc.Close()

	r := &pb.AddReq{
		UserUUID:   req.UserUUID,
		DeviceUUID: req.DeviceUUID,
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	_, err = cl.AddToFavourite(ctx, r)
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

func (h *handler) RemoveFromFavourite(c *gin.Context) {
	var req models.RemoveReq

	if err := c.ShouldBindJSON(&req); err != nil {
		responser.UserError(c.Writer, "failed to decode req body")
		return
	}

	if err := models.Validate(&req); err != nil {
		responser.UserError(c.Writer, err.Error())
		return
	}

	cl, cc, err := client.DialUser(USER_ADDR)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}
	defer cc.Close()

	r := &pb.RemoveReq{
		UserUUID:   req.UserUUID,
		DeviceUUID: req.DeviceUUID,
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	_, err = cl.RemoveFromFavourite(ctx, r)
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

func (h *handler) GetFavourite(c *gin.Context) {
	userUUID := c.Param("userUUID")

	if userUUID == "" {
		responser.UserError(c.Writer, "incorrect URL")
		return
	}

	cl, cc, err := client.DialUser(USER_ADDR)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}
	defer cc.Close()

	r := &pb.GetReq{
		UserUUID: userUUID,
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	coll, err := cl.GetFavourite(ctx, r)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			responser.UserError(c.Writer, st.Message())
			return
		}
		responser.ServerError(c.Writer, err)
		return
	}

	responser.Data(c.Writer, responser.H{
		"amount": len(coll.Devices),
		"data":   coll.Devices,
	})
}

func (h *handler) AddToCart(c *gin.Context) {
	var req models.AddReq

	if err := c.ShouldBindJSON(&req); err != nil {
		responser.UserError(c.Writer, "failed to decode req body")
		return
	}

	if err := models.Validate(&req); err != nil {
		responser.UserError(c.Writer, err.Error())
		return
	}

	cl, cc, err := client.DialUser(USER_ADDR)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}
	defer cc.Close()

	r := &pb.AddReq{
		UserUUID:   req.UserUUID,
		DeviceUUID: req.DeviceUUID,
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	_, err = cl.AddToCart(ctx, r)
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

func (h *handler) RemoveFromCart(c *gin.Context) {
	var req models.RemoveReq

	if err := c.ShouldBindJSON(&req); err != nil {
		responser.UserError(c.Writer, "failed to decode req body")
		return
	}

	if err := models.Validate(&req); err != nil {
		responser.UserError(c.Writer, err.Error())
		return
	}

	cl, cc, err := client.DialUser(USER_ADDR)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}
	defer cc.Close()

	r := &pb.RemoveReq{
		UserUUID:   req.UserUUID,
		DeviceUUID: req.DeviceUUID,
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	_, err = cl.RemoveFromCart(ctx, r)
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

func (h *handler) GetCart(c *gin.Context) {
	userUUID := c.Param("userUUID")

	if userUUID == "" {
		responser.UserError(c.Writer, "incorrect URL")
		return
	}

	cl, cc, err := client.DialUser(USER_ADDR)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}
	defer cc.Close()

	r := &pb.GetReq{
		UserUUID: userUUID,
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	coll, err := cl.GetCart(ctx, r)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			responser.UserError(c.Writer, st.Message())
			return
		}
		responser.ServerError(c.Writer, err)
		return
	}

	responser.Data(c.Writer, responser.H{
		"amount": len(coll.Devices),
		"data":   coll.Devices,
	})
}

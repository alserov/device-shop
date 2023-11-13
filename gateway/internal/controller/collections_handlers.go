package controller

import (
	"context"
	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/gateway/pkg/models"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	"github.com/alserov/device-shop/proto/gen"
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
	msg, err := utils.RequestToPBMessage[models.AddToCollectionReq, pb.AddReq](c.Request, utils.AddReqToPB)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}

	cl, cc, err := client.DialUser(USER_ADDR)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	_, err = cl.AddToFavourite(ctx, msg)
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
	msg, err := utils.RequestToPBMessage[models.RemoveDeviceReq, pb.RemoveReq](c.Request, utils.RemoveReqToPB)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}

	cl, cc, err := client.DialUser(USER_ADDR)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	_, err = cl.RemoveFromFavourite(ctx, msg)
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

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	coll, err := cl.GetFavourite(ctx, &pb.GetReq{UserUUID: userUUID})
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
	msg, err := utils.RequestToPBMessage[models.AddToCollectionReq, pb.AddReq](c.Request, utils.AddReqToPB)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}

	cl, cc, err := client.DialUser(USER_ADDR)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	_, err = cl.AddToCart(ctx, msg)
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
	msg, err := utils.RequestToPBMessage[models.RemoveDeviceReq, pb.RemoveReq](c.Request, utils.RemoveReqToPB)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}
	cl, cc, err := client.DialUser(USER_ADDR)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	_, err = cl.RemoveFromCart(ctx, msg)
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

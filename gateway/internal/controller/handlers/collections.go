package handlers

import (
	"context"
	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	"github.com/alserov/device-shop/proto/gen"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
	"net/http"
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

type collectionsHandler struct {
	userAddr string
	logger   *logrus.Logger
}

func NewCollectionsHandler(userAddr string, logger *logrus.Logger) CollectionsHandler {
	return &collectionsHandler{
		userAddr: userAddr,
		logger:   logger,
	}
}

func (h *collectionsHandler) AddToFavourite(c *gin.Context) {
	addCred, err := utils.Decode[pb.ChangeCollectionReq](c.Request, utils.CheckCollection)
	if err != nil {
		responser.UserError(c.Writer, err.Error())
		return
	}

	cl, cc, err := client.DialUser(h.userAddr)
	if err != nil {
		responser.ServerError(c.Writer, h.logger, err)
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	_, err = cl.AddToFavourite(ctx, addCred)
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

func (h *collectionsHandler) RemoveFromFavourite(c *gin.Context) {
	removeCred, err := utils.Decode[pb.ChangeCollectionReq](c.Request, utils.CheckCollection)
	if err != nil {
		responser.UserError(c.Writer, err.Error())
		return
	}

	cl, cc, err := client.DialUser(h.userAddr)
	if err != nil {
		responser.ServerError(c.Writer, h.logger, err)
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	_, err = cl.RemoveFromFavourite(ctx, removeCred)
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

func (h *collectionsHandler) GetFavourite(c *gin.Context) {
	userUUID := c.Param("userUUID")

	if userUUID == "" {
		responser.UserError(c.Writer, "incorrect URL")
		return
	}

	cl, cc, err := client.DialUser(h.userAddr)
	if err != nil {
		responser.ServerError(c.Writer, h.logger, err)
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	coll, err := cl.GetFavourite(ctx, &pb.GetCollectionReq{UserUUID: userUUID})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			responser.UserError(c.Writer, st.Message())
			return
		}
		responser.ServerError(c.Writer, h.logger, err)
		return
	}

	responser.Data(c.Writer, responser.H{
		"amount": len(coll.Devices),
		"data":   coll.Devices,
	})
}

func (h *collectionsHandler) AddToCart(c *gin.Context) {
	addCred, err := utils.Decode[pb.ChangeCollectionReq](c.Request, utils.CheckCollection)
	if err != nil {
		responser.UserError(c.Writer, err.Error())
		return
	}

	cl, cc, err := client.DialUser(h.userAddr)
	if err != nil {
		responser.ServerError(c.Writer, h.logger, err)
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	_, err = cl.AddToCart(ctx, addCred)
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

func (h *collectionsHandler) RemoveFromCart(c *gin.Context) {
	removeCred, err := utils.Decode[pb.ChangeCollectionReq](c.Request, utils.CheckCollection)
	if err != nil {
		responser.UserError(c.Writer, err.Error())
		return
	}

	cl, cc, err := client.DialUser(h.userAddr)
	if err != nil {
		responser.ServerError(c.Writer, h.logger, err)
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	_, err = cl.RemoveFromCart(ctx, removeCred)
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

func (h *collectionsHandler) GetCart(c *gin.Context) {
	userUUID := c.Param("userUUID")

	if userUUID == "" {
		responser.UserError(c.Writer, "incorrect URL")
		return
	}

	cl, cc, err := client.DialUser(h.userAddr)
	if err != nil {
		responser.ServerError(c.Writer, h.logger, err)
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	coll, err := cl.GetCart(ctx, &pb.GetCollectionReq{
		UserUUID: userUUID,
	})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			responser.UserError(c.Writer, st.Message())
			return
		}
		responser.ServerError(c.Writer, h.logger, err)
		return
	}

	responser.Data(c.Writer, responser.H{
		"amount": len(coll.Devices),
		"data":   coll.Devices,
	})
}

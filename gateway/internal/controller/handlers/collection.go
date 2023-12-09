package handlers

import (
	"context"
	"github.com/alserov/device-shop/gateway/internal/controller/handlers/models"
	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/internal/utils/validation"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	"github.com/alserov/device-shop/proto/gen/collection"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
	"log/slog"
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
	services models.Services
	log      *slog.Logger
}

func NewCollectionsHandler(userAddr string, log *slog.Logger) CollectionsHandler {
	return &collectionsHandler{
		services: models.Services{
			User: models.Service{
				Addr: userAddr,
			},
		},
		log: log,
	}
}

func (h *collectionsHandler) AddToFavourite(c *gin.Context) {
	addCred, err := utils.Decode[collection.ChangeCollectionReq](c.Request, validation.CheckCollection)
	if err != nil {
		responser.UserError(c.Writer, err.Error())
		return
	}

	cl, cc, err := client.DialCollection(h.services.User.Addr)
	if err != nil {
		responser.ServerError(c.Writer, h.log, err)
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
		responser.ServerError(c.Writer, h.log, err)
		return
	}

	c.Status(http.StatusOK)
}

func (h *collectionsHandler) RemoveFromFavourite(c *gin.Context) {
	removeCred, err := utils.Decode[collection.ChangeCollectionReq](c.Request, validation.CheckCollection)
	if err != nil {
		responser.UserError(c.Writer, err.Error())
		return
	}

	cl, cc, err := client.DialCollection(h.services.User.Addr)
	if err != nil {
		responser.ServerError(c.Writer, h.log, err)
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
		responser.ServerError(c.Writer, h.log, err)
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

	cl, cc, err := client.DialCollection(h.services.User.Addr)
	if err != nil {
		responser.ServerError(c.Writer, h.log, err)
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	coll, err := cl.GetFavourite(ctx, &collection.GetCollectionReq{UserUUID: userUUID})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			responser.UserError(c.Writer, st.Message())
			return
		}
		responser.ServerError(c.Writer, h.log, err)
		return
	}

	responser.Data(c.Writer, responser.H{
		"amount": len(coll.Devices),
		"data":   coll.Devices,
	})
}

func (h *collectionsHandler) AddToCart(c *gin.Context) {
	addCred, err := utils.Decode[collection.ChangeCollectionReq](c.Request, validation.CheckCollection)
	if err != nil {
		responser.UserError(c.Writer, err.Error())
		return
	}

	cl, cc, err := client.DialCollection(h.services.User.Addr)
	if err != nil {
		responser.ServerError(c.Writer, h.log, err)
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
		responser.ServerError(c.Writer, h.log, err)
		return
	}

	c.Status(http.StatusOK)
}

func (h *collectionsHandler) RemoveFromCart(c *gin.Context) {
	removeCred, err := utils.Decode[collection.ChangeCollectionReq](c.Request, validation.CheckCollection)
	if err != nil {
		responser.UserError(c.Writer, err.Error())
		return
	}

	cl, cc, err := client.DialCollection(h.services.User.Addr)
	if err != nil {
		responser.ServerError(c.Writer, h.log, err)
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
		responser.ServerError(c.Writer, h.log, err)
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

	cl, cc, err := client.DialCollection(h.services.User.Addr)
	if err != nil {
		responser.ServerError(c.Writer, h.log, err)
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	coll, err := cl.GetCart(ctx, &collection.GetCollectionReq{
		UserUUID: userUUID,
	})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			responser.UserError(c.Writer, st.Message())
			return
		}
		responser.ServerError(c.Writer, h.log, err)
		return
	}

	responser.Data(c.Writer, responser.H{
		"amount": len(coll.Devices),
		"data":   coll.Devices,
	})
}

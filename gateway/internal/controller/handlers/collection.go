package handlers

import (
	"context"
	"github.com/alserov/device-shop/gateway/internal/logger"
	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/internal/utils/validation"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	"github.com/alserov/device-shop/proto/gen/collection"
	"github.com/gin-gonic/gin"
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
	serviceAddr string
	log         *slog.Logger
}

type CollectionH struct {
	UserAddr string
	Log      *slog.Logger
}

func NewCollectionsHandler(ch *CollectionH) CollectionsHandler {
	return &collectionsHandler{
		serviceAddr: ch.UserAddr,
		log:         ch.Log,
	}
}

func (h *collectionsHandler) AddToFavourite(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "collectionsHandler.AddToFavourite"

	addCred, err := utils.Decode[collection.ChangeCollectionReq](c.Request, validation.CheckCollection)
	if err != nil {
		w.UserError(err.Error())
		return
	}

	cl, cc, err := client.DialCollection(h.serviceAddr)
	if err != nil {
		h.log.Error("failed to dial collection service", logger.Error(err, op))
		w.ServerError()
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	_, err = cl.AddToFavourite(ctx, addCred)
	if err != nil {
		w.HandleServiceError(err, "cl.AddToFavourite", h.log)
		return
	}

	w.OK()
}

func (h *collectionsHandler) RemoveFromFavourite(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "collectionsHandler.RemoveFromFavourite"

	deviceAndUserUUIDs, err := utils.Decode[collection.ChangeCollectionReq](c.Request, validation.CheckCollection)
	if err != nil {
		w.UserError(err.Error())
		return
	}

	cl, cc, err := client.DialCollection(h.serviceAddr)
	if err != nil {
		h.log.Error("failed to dial collection service", logger.Error(err, op))
		w.ServerError()
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	_, err = cl.RemoveFromFavourite(ctx, deviceAndUserUUIDs)
	if err != nil {
		w.HandleServiceError(err, "cl.AddToFavourite", h.log)
		return
	}

	w.OK()
}

func (h *collectionsHandler) GetFavourite(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "collectionsHandler.GetFavourite"

	userUUID := c.Param("userUUID")

	if userUUID == "" {
		w.UserError(invalidQueryParam)
		return
	}

	cl, cc, err := client.DialCollection(h.serviceAddr)
	if err != nil {
		h.log.Error("failed to dial collection service", logger.Error(err, op))
		w.ServerError()
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	coll, err := cl.GetFavourite(ctx, &collection.GetCollectionReq{UserUUID: userUUID})
	if err != nil {
		w.HandleServiceError(err, "cl.GetFavourite", h.log)
		return
	}

	w.Data(responser.H{
		"amount":    len(coll.Devices),
		"favourite": coll.Devices,
	})
}

func (h *collectionsHandler) AddToCart(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "collectionsHandler.AddToCart"

	addCred, err := utils.Decode[collection.ChangeCollectionReq](c.Request, validation.CheckCollection)
	if err != nil {
		w.UserError(err.Error())
		return
	}

	cl, cc, err := client.DialCollection(h.serviceAddr)
	if err != nil {
		h.log.Error("failed to dial collection service", logger.Error(err, op))
		w.ServerError()
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	_, err = cl.AddToCart(ctx, addCred)
	if err != nil {
		w.HandleServiceError(err, "cl.AddToCart", h.log)
		return
	}

	w.OK()
}

func (h *collectionsHandler) RemoveFromCart(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "collectionsHandler.RemoveFromCart"

	removeCred, err := utils.Decode[collection.ChangeCollectionReq](c.Request, validation.CheckCollection)
	if err != nil {
		w.UserError(err.Error())
		return
	}

	cl, cc, err := client.DialCollection(h.serviceAddr)
	if err != nil {
		h.log.Error("failed to dial collection service", logger.Error(err, op))
		w.ServerError()
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	_, err = cl.RemoveFromCart(ctx, removeCred)
	if err != nil {
		w.HandleServiceError(err, "cl.RemoveFromCart", h.log)
		return
	}

	c.Status(http.StatusOK)
}

func (h *collectionsHandler) GetCart(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "collectionsHandler.GetCart"

	userUUID := c.Param("userUUID")

	if userUUID == "" {
		w.UserError(invalidQueryParam)
		return
	}

	cl, cc, err := client.DialCollection(h.serviceAddr)
	if err != nil {
		h.log.Error("failed to dial collection service", logger.Error(err, op))
		w.ServerError()
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	coll, err := cl.GetCart(ctx, &collection.GetCollectionReq{
		UserUUID: userUUID,
	})
	if err != nil {
		w.HandleServiceError(err, "cl.RemoveFromCart", h.log)
		return
	}

	w.Data(responser.H{
		"amount": len(coll.Devices),
		"data":   coll.Devices,
	})
}

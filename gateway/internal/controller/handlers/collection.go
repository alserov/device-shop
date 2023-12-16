package handlers

import (
	"context"
	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/internal/utils/validation"
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
	client collection.CollectionsClient
	log    *slog.Logger
}

func NewCollectionsHandler(c collection.CollectionsClient, log *slog.Logger) CollectionsHandler {
	return &collectionsHandler{
		client: c,
		log:    log,
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

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	_, err = h.client.AddToFavourite(ctx, addCred)
	if err != nil {
		w.HandleServiceError(err, op, h.log)
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
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	_, err = h.client.RemoveFromFavourite(ctx, deviceAndUserUUIDs)
	if err != nil {
		w.HandleServiceError(err, op, h.log)
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

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	coll, err := h.client.GetFavourite(ctx, &collection.GetCollectionReq{UserUUID: userUUID})
	if err != nil {
		w.HandleServiceError(err, op, h.log)
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

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	_, err = h.client.AddToCart(ctx, addCred)
	if err != nil {
		w.HandleServiceError(err, op, h.log)
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

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	_, err = h.client.RemoveFromCart(ctx, removeCred)
	if err != nil {
		w.HandleServiceError(err, op, h.log)
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

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	coll, err := h.client.GetCart(ctx, &collection.GetCollectionReq{
		UserUUID: userUUID,
	})
	if err != nil {
		w.HandleServiceError(err, op, h.log)
		return
	}

	w.Data(responser.H{
		"amount": len(coll.Devices),
		"data":   coll.Devices,
	})
}

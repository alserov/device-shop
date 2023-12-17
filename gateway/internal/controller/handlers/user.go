package handlers

import (
	"context"

	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/internal/utils/validation"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	"github.com/alserov/device-shop/proto/gen/user"

	"github.com/gin-gonic/gin"
	"log/slog"
	"time"
)

type UserHandler interface {
	TopUpBalance(ctx *gin.Context)
	GetInfo(c *gin.Context)
}

type userHandler struct {
	client user.UsersClient
	log    *slog.Logger
}

func NewUserHandler(c user.UsersClient, log *slog.Logger) UserHandler {
	return &userHandler{
		client: c,
		log:    log,
	}
}

func (h *userHandler) TopUpBalance(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "userHandler.TopUpBalance"

	cashAmount, err := utils.Decode[user.BalanceReq](c.Request, validation.CheckTopUpBalance)
	if err != nil {
		w.UserError(err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	res, err := h.client.TopUpBalance(ctx, cashAmount)
	if err != nil {
		w.HandleServiceError(err, op, h.log)
		return
	}

	w.Data(responser.H{
		"cash": res.Cash,
	})
}

func (h *userHandler) GetInfo(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "authHandler.getInfo"

	userUUID := c.Param("user_uuid")

	if userUUID == "" {
		w.UserError("userUUID can not be empty")
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	res, err := h.client.GetUserInfo(ctx, &user.GetUserInfoReq{
		UserUUID: userUUID,
	})
	if err != nil {
		w.HandleServiceError(err, op, h.log)
		return
	}

	w.Value(res)
}

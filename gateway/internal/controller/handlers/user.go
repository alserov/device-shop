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

type UsersHandler interface {
	TopUpBalance(ctx *gin.Context)
}

type usersHandler struct {
	client user.UsersClient
	log    *slog.Logger
}

func NewUserHandler(c user.UsersClient, log *slog.Logger) UsersHandler {
	return &usersHandler{
		client: c,
		log:    log,
	}
}

func (h *usersHandler) TopUpBalance(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "usersHandler.TopUpBalance"

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

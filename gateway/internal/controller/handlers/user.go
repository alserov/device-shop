package handlers

import (
	"context"
	"github.com/alserov/device-shop/gateway/internal/controller/handlers/models"
	"github.com/alserov/device-shop/gateway/internal/logger"
	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/internal/utils/validation"
	"github.com/alserov/device-shop/gateway/pkg/client"
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
	services models.Services
	log      *slog.Logger
}

func NewUserHandler(userAddr string, log *slog.Logger) UsersHandler {
	return &usersHandler{
		services: models.Services{
			User: models.Service{
				Addr: userAddr,
			},
		},
		log: log,
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

	cl, cc, err := client.DialUser(h.services.User.Addr)
	if err != nil {
		h.log.Error("failed to dial user service", logger.Error(err, op))
		w.ServerError()
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	res, err := cl.TopUpBalance(ctx, cashAmount)
	if err != nil {
		w.HandleServiceError(err, "cl.TopUpBalance", h.log)
		return
	}

	w.Data(responser.H{
		"cash": res.Cash,
	})
}

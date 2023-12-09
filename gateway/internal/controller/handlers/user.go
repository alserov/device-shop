package handlers

import (
	"context"
	"github.com/alserov/device-shop/gateway/internal/controller/handlers/models"
	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/internal/utils/validation"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	"github.com/alserov/device-shop/proto/gen/user"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
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
	cashAmount, err := utils.Decode[user.BalanceReq](c.Request, validation.CheckTopUpBalance)
	if err != nil {
		responser.UserError(c.Writer, err.Error())
		return
	}

	cl, cc, err := client.DialUser(h.services.User.Addr)
	if err != nil {
		responser.ServerError(c.Writer, h.log, err)
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	res, err := cl.TopUpBalance(ctx, cashAmount)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			responser.UserError(c.Writer, st.Message())
			return
		}
		responser.ServerError(c.Writer, h.log, err)
		return
	}

	responser.Data(c.Writer, responser.H{
		"cash": res.Cash,
	})
}

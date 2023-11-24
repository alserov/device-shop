package handlers

import (
	"context"
	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	pb "github.com/alserov/device-shop/proto/gen"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
	"time"
)

type UsersHandler interface {
	TopUpBalance(ctx *gin.Context)
	GetInfo(ctx *gin.Context)
}

type usersHandler struct {
	userAddr string
	logger   *logrus.Logger
}

func NewUserHandler(userAddr string, logger *logrus.Logger) UsersHandler {
	return &usersHandler{
		userAddr: userAddr,
		logger:   logger,
	}
}

func (h *usersHandler) TopUpBalance(c *gin.Context) {
	cashAmount, err := utils.Decode[pb.BalanceReq](c.Request, utils.CheckTopUpBalance)
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

	res, err := cl.TopUpBalance(ctx, cashAmount)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			responser.UserError(c.Writer, st.Message())
			return
		}
		responser.ServerError(c.Writer, h.logger, err)
		return
	}

	responser.Data(c.Writer, responser.H{
		"cash": res.Cash,
	})
}

func (h *usersHandler) GetInfo(c *gin.Context) {
	userUUID := c.Param("userUUID")

	if userUUID == "" {
		responser.UserError(c.Writer, "userUUID cannot be empty")
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

	res, err := cl.GetUserInfo(ctx, &pb.GetUserInfoReq{
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

	responser.Value(c.Writer, res)
}

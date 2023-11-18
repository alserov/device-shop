package controller

import (
	"context"
	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	pb "github.com/alserov/device-shop/proto/gen"
	"github.com/alserov/device-shop/user-service/pkg/entity"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
	"time"
)

type Balancer interface {
	TopUpBalance(ctx *gin.Context)
}

func (h *handler) TopUpBalance(c *gin.Context) {
	msg, err := utils.RequestToPBMessage[entity.TopUpBalanceReq, pb.TopUpBalanceReq](c.Request, utils.TopUpBalanceReqToPB)
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

	res, err := cl.TopUpBalance(ctx, msg)
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

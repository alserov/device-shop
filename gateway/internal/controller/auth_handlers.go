package controller

import (
	"context"
	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	"github.com/alserov/device-shop/proto/gen"
	user "github.com/alserov/device-shop/user-service/pkg/entity"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
	"time"
)

type Auther interface {
	Signup(c *gin.Context)
	Login(c *gin.Context)
}

func (h *handler) Signup(c *gin.Context) {
	msg, err := utils.RequestToPBMessage[user.SignupReq, pb.SignupReq](c.Request, utils.SignupReqToPB)
	if err != nil {
		responser.ServerError(c.Writer, h.logger, err)
		return
	}

	if valid := govalidator.IsEmail(msg.Email); !valid {
		responser.UserError(c.Writer, "invalid email")
		return
	}

	cl, cc, err := client.DialUser(h.userAddr)
	if err != nil {
		responser.ServerError(c.Writer, h.logger, err)
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	user, err := cl.Signup(ctx, msg)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			responser.UserError(c.Writer, st.Message())
			return
		}
		responser.ServerError(c.Writer, h.logger, err)
		return
	}

	c.SetCookie("token", user.Token, 604800, "/", "localhost", false, true)
	responser.Value(c.Writer, user)
}

func (h *handler) Login(c *gin.Context) {
	msg, err := utils.RequestToPBMessage[user.LoginReq, pb.LoginReq](c.Request, utils.LoginReqToPB)
	if err != nil {
		responser.ServerError(c.Writer, h.logger, err)
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

	res, err := cl.Login(ctx, msg)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			responser.UserError(c.Writer, st.Message())
			return
		}
		responser.ServerError(c.Writer, h.logger, err)
		return
	}

	c.SetCookie("token", res.Token, 604800, "/", "localhost", false, true)

	responser.Data(c.Writer, responser.H{
		"refreshToken": res.RefreshToken,
		"userUUID":     res.UUID,
	})
}

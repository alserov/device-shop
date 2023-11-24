package handlers

import (
	"context"
	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	"github.com/alserov/device-shop/proto/gen"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
	"time"
)

type AuthHandler interface {
	Signup(c *gin.Context)
	Login(c *gin.Context)
}

type authHandler struct {
	logger   *logrus.Logger
	authAddr string
}

func NewAuthHandler(authAddr string, logger *logrus.Logger) AuthHandler {
	return &authHandler{
		logger:   logger,
		authAddr: authAddr,
	}
}

func (h *authHandler) Signup(c *gin.Context) {
	userInfo, err := utils.Decode[pb.SignupReq](c.Request, utils.CheckSignup)
	if err != nil {
		responser.UserError(c.Writer, err.Error())
		return
	}

	if valid := govalidator.IsEmail(userInfo.Email); !valid {
		responser.UserError(c.Writer, "invalid email")
		return
	}

	cl, cc, err := client.DialAuth(h.authAddr)
	if err != nil {
		responser.ServerError(c.Writer, h.logger, err)
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	user, err := cl.Signup(ctx, userInfo)
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

func (h *authHandler) Login(c *gin.Context) {
	userInfo, err := utils.Decode[pb.LoginReq](c.Request, utils.CheckLogin)
	if err != nil {
		responser.UserError(c.Writer, err.Error())
		return
	}

	cl, cc, err := client.DialAuth(h.authAddr)
	if err != nil {
		responser.ServerError(c.Writer, h.logger, err)
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	res, err := cl.Login(ctx, userInfo)
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

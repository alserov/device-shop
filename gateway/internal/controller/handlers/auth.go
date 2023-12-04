package handlers

import (
	"context"
	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/internal/utils/validation"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	"github.com/alserov/device-shop/proto/gen/auth"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"time"
)

type AuthHandler interface {
	Signup(c *gin.Context)
	Login(c *gin.Context)
	GetInfo(c *gin.Context)
}

type authHandler struct {
	logger   *slog.Logger
	authAddr string
}

func NewAuthHandler(authAddr string, logger *slog.Logger) AuthHandler {
	return &authHandler{
		logger:   logger,
		authAddr: authAddr,
	}
}

func (h *authHandler) Signup(c *gin.Context) {
	userInfo, err := utils.Decode[auth.SignupReq](c.Request, validation.CheckSignup)
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
			switch st.Code() {
			case codes.Internal:
				responser.ServerError(c.Writer, h.logger, err)
			default:
				responser.UserError(c.Writer, st.Message())
			}
			return
		}
	}

	c.SetCookie("token", user.Token, 604800, "/", "localhost", false, true)
	responser.Value(c.Writer, user)
}

func (h *authHandler) Login(c *gin.Context) {
	userInfo, err := utils.Decode[auth.LoginReq](c.Request, validation.CheckLogin)
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
			switch st.Code() {
			case codes.Internal:
				responser.ServerError(c.Writer, h.logger, err)
			default:
				responser.UserError(c.Writer, st.Message())
			}
			return
		}
	}

	c.SetCookie("token", res.Token, 604800, "/", "localhost", false, true)

	responser.Data(c.Writer, responser.H{
		"refreshToken": res.RefreshToken,
		"userUUID":     res.UUID,
	})
}

func (h *authHandler) GetInfo(c *gin.Context) {
	userUUID := c.Param("userUUID")

	if userUUID == "" {
		responser.UserError(c.Writer, "userUUID cannot be empty")
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

	res, err := cl.GetUserInfo(ctx, &auth.GetUserInfoReq{
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
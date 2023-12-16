package handlers

import (
	"context"
	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/internal/utils/validation"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	"github.com/alserov/device-shop/proto/gen/user"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"log/slog"
	"time"
)

type AuthHandler interface {
	Signup(c *gin.Context)
	Login(c *gin.Context)
	GetInfo(c *gin.Context)
}

type authHandler struct {
	log    *slog.Logger
	client user.UsersClient
}

func NewAuthHandler(c user.UsersClient, log *slog.Logger) AuthHandler {
	return &authHandler{
		log:    log,
		client: c,
	}
}

func (h *authHandler) Signup(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "authHandler.Signup"

	userInfo, err := utils.Decode[user.SignupReq](c.Request, validation.CheckSignup)
	if err != nil {
		w.UserError(err.Error())
		return
	}

	if valid := govalidator.IsEmail(userInfo.Email); !valid {
		w.UserError("invalid email")
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	user, err := h.client.Signup(ctx, userInfo)
	if err != nil {
		w.HandleServiceError(err, op, h.log)
		return
	}

	c.SetCookie("token", user.Token, 604800, "/", "localhost", false, true)

	w.Value(user)
}

func (h *authHandler) Login(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "authHandler.Login"

	userInfo, err := utils.Decode[user.LoginReq](c.Request, validation.CheckLogin)
	if err != nil {
		w.UserError(err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	res, err := h.client.Login(ctx, userInfo)
	if err != nil {
		w.HandleServiceError(err, op, h.log)
		return
	}

	c.SetCookie("token", res.Token, 604800, "/", "localhost", false, true)

	w.Data(responser.H{
		"refreshToken": res.RefreshToken,
		"userUUID":     res.UUID,
	})
}

func (h *authHandler) GetInfo(c *gin.Context) {
	w := responser.NewResponser(c.Writer)
	op := "authHandler.getInfo"

	userUUID := c.Param("userUUID")

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

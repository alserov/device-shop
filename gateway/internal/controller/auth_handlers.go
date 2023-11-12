package controller

import (
	"context"
	"github.com/alserov/shop/api/pkg/client"
	"github.com/alserov/shop/api/pkg/responser"
	"github.com/alserov/shop/gateway/pkg/models"
	"github.com/alserov/shop/user-service/pkg/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
	"time"
)

type AuthHandler interface {
	Signup(c *gin.Context)
	Login(c *gin.Context)
}

// @Summary Signup
// @Tags auth
// @Description Creates new user
// @ID create-user
// @Accept json
// @Produce json
// @Param input body models.SignupReq true "user credentials"
// @Success 200
// @Failure 400 {object} responser.WithError
// @Failure 500
// @Router /auth/signup [post]

func (h *handler) Signup(c *gin.Context) {
	var req models.SignupReq

	if err := c.ShouldBindJSON(&req); err != nil {
		responser.UserError(c.Writer, "failed to parse req body")
		return
	}

	if invalidEmail := req.Validate(); invalidEmail != nil {
		responser.UserError(c.Writer, invalidEmail.Error())
		return
	}

	cl, cc, err := client.DialUsers(USER_ADDR)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}
	defer cc.Close()

	r := &pb.SignupReq{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	user, err := cl.Signup(ctx, r)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			responser.UserError(c.Writer, st.Message())
			return
		}
		responser.ServerError(c.Writer, err)
		return
	}

	c.SetCookie("token", user.Token, 604800, "/", "localhost", false, true)
	responser.Value(c.Writer, user)
}

func (h *handler) Login(c *gin.Context) {
	var req models.LoginReq

	if err := c.ShouldBindJSON(&req); err != nil {
		responser.UserError(c.Writer, "failed to parse req body")
		return
	}

	if err := models.Validate(&req); err != nil {
		responser.UserError(c.Writer, err.Error())
		return
	}

	cl, cc, err := client.DialUsers(USER_ADDR)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}
	defer cc.Close()

	r := &pb.LoginReq{
		Username: req.Username,
		Password: req.Password,
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	tokens, err := cl.Login(ctx, r)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			responser.UserError(c.Writer, st.Message())
			return
		}
		responser.ServerError(c.Writer, err)
		return
	}

	c.SetCookie("token", tokens.Token, 604800, "/", "localhost", false, true)

	responser.Data(c.Writer, responser.H{
		"refreshToken": tokens.RefreshToken,
	})
}

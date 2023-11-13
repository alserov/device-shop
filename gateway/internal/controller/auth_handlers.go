package controller

import (
	"context"
	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/pkg/client"
	"github.com/alserov/device-shop/gateway/pkg/models"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	"github.com/alserov/device-shop/proto/gen"
	"github.com/asaskevich/govalidator"
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
	msg, err := utils.RequestToPBMessage[models.SignupReq, pb.SignupReq](c.Request, utils.SignupReqToPB)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}

	if valid := govalidator.IsEmail(msg.Email); !valid {
		responser.UserError(c.Writer, "invalid email")
		return
	}

	cl, cc, err := client.DialUser(USER_ADDR)
	if err != nil {
		responser.ServerError(c.Writer, err)
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
		responser.ServerError(c.Writer, err)
		return
	}

	c.SetCookie("token", user.Token, 604800, "/", "localhost", false, true)
	responser.Value(c.Writer, user)
}

func (h *handler) Login(c *gin.Context) {
	msg, err := utils.RequestToPBMessage[models.LoginReq, pb.LoginReq](c.Request, utils.LoginReqToPB)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}

	cl, cc, err := client.DialUser(USER_ADDR)
	if err != nil {
		responser.ServerError(c.Writer, err)
		return
	}
	defer cc.Close()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	tokens, err := cl.Login(ctx, msg)
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

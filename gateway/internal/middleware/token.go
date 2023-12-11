package middleware

import (
	"github.com/alserov/device-shop/gateway/internal/utils"
	"github.com/alserov/device-shop/gateway/pkg/responser"
	"github.com/gin-gonic/gin"
)

func CheckIfAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := responser.NewResponser(c.Writer)

		token, err := c.Cookie("token")
		if err != nil {
			w.UserError(err.Error())
			c.Abort()
		}

		if token == "" {
			w.UserError("not authorized")
			c.Abort()
		}

		if err = utils.ValidateToken(token); err != nil {
			w.UserError(err.Error())
			c.Abort()
		}

		c.Next()
	}
}

func CheckIfAllowed() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := responser.NewResponser(c.Writer)

		token, err := c.Cookie("token")
		if err != nil {
			w.UserError(err.Error())
			c.Abort()
			return
		}

		if token == "" {
			w.UserError("not authorized")
			c.Abort()
			return
		}

		if err = utils.CheckIfAdmin(token); err != nil {
			w.NotAllowed()
			c.Abort()
			return
		}

		c.Next()
	}
}

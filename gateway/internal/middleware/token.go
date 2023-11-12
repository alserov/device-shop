package middleware

import (
	"github.com/alserov/shop/gateway/internal/utils"
	"github.com/alserov/shop/gateway/pkg/responser"
	"github.com/gin-gonic/gin"
)

func CheckIfAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			responser.ServerError(c.Writer, err)
			c.Abort()
		}

		if token == "" {
			responser.UserError(c.Writer, "not authorized")
			c.Abort()
		}

		if err = utils.ValidateToken(token); err != nil {
			responser.UserError(c.Writer, err.Error())
			c.Abort()
		}

		c.Next()
	}
}

func CheckIfAllowed() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			responser.UserError(c.Writer, err.Error())
			c.Abort()
			return
		}

		if token == "" {
			responser.UserError(c.Writer, "not authorized")
			c.Abort()
			return
		}

		if err = utils.CheckIfAdmin(token); err != nil {
			responser.NotAllowed(c.Writer)
			c.Abort()
			return
		}

		c.Next()
	}
}

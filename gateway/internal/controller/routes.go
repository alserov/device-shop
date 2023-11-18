package controller

import (
	"github.com/alserov/device-shop/gateway/internal/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	authPath    = "/auth"
	adminPath   = "/admin"
	devicesPath = "/devices"
	actionsPath = "/actions"
	ordersPath  = "/orders"
)

func LoadRoutes(r *gin.Engine, h Handler) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// AUTH
	userAuth := r.Group(authPath)
	userAuth.POST("/signup", h.Signup)
	userAuth.POST("/login", h.Login)

	// USER ACTIONS
	userActions := r.Group(actionsPath).Use(middleware.CheckIfAuthorized())
	userActions.PUT("/balance", h.TopUpBalance)
	userActions.POST("/new-favourite", h.AddToFavourite)
	userActions.DELETE("/delete-favourite", h.RemoveFromFavourite)
	userActions.GET("/favourite/:userUUID", h.GetFavourite)
	userActions.POST("/new-cart", h.AddToCart)
	userActions.DELETE("/delete-cart", h.RemoveFromCart)
	userActions.GET("/cart/:userUUID", h.GetCart)
	userActions.GET("/info/:userUUID", h.GetInfo)

	// ORDERS
	order := r.Group(ordersPath).Use(middleware.CheckIfAuthorized())
	order.POST("/new", h.CreateOrder)
	order.PUT("/update", h.UpdateOrder)
	order.GET("/:orderUUID", h.CheckOrder)

	// ADMIN routes
	admin := r.Group(adminPath).Use(middleware.CheckIfAllowed())
	admin.POST("/new-device", h.CreateDevice)
	admin.DELETE("/delete-device/:deviceUUID", h.DeleteDevice)
	admin.PUT("/update-device", h.UpdateDevice)

	// DEVICE ROUTES
	device := r.Group(devicesPath)
	device.GET("/", h.GetAllDevices)
	device.GET("/title/:title", h.GetDevicesByTitle)
	device.GET("/manufacturer/:manu", h.GetDevicesByManufacturer)
	device.GET("/price", h.GetDevicesByPrice)
}

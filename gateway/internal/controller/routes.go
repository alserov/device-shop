package controller

import (
	"github.com/alserov/device-shop/gateway/internal/middleware"
	"github.com/gin-gonic/gin"
)

const (
	authPath    = "/auth"
	adminPath   = "/admin"
	devicesPath = "/devices"
	userPath    = "/user"
	ordersPath  = "/orders"
)

func LoadRoutes(r *gin.Engine, h *Controller) {
	// AUTH
	userAuth := r.Group(authPath)
	userAuth.POST("/signup", h.authHandler.Signup)
	userAuth.POST("/login", h.authHandler.Login)

	// USER ACTIONS
	userActions := r.Group(userPath).Use(middleware.CheckIfAuthorized())
	userActions.GET("/info/:user_uuid", h.userHandler.GetInfo)
	userActions.PUT("/balance", h.userHandler.TopUpBalance)
	userActions.POST("/new-favourite", h.collectionHandler.AddToFavourite)
	userActions.DELETE("/delete-favourite", h.collectionHandler.RemoveFromFavourite)
	userActions.GET("/favourite/:userUUID", h.collectionHandler.GetFavourite)
	userActions.POST("/new-cart", h.collectionHandler.AddToCart)
	userActions.DELETE("/delete-cart", h.collectionHandler.RemoveFromCart)
	userActions.GET("/cart/:userUUID", h.collectionHandler.GetCart)

	// ORDERS
	order := r.Group(ordersPath).Use(middleware.CheckIfAuthorized())
	order.POST("/new", h.orderHandler.CreateOrder)
	order.PUT("/update", h.orderHandler.UpdateOrder)
	order.GET("/:order_uuid", h.orderHandler.CheckOrder)
	order.DELETE("/:order_uuid", h.orderHandler.CancelOrder)

	// ADMIN routes
	admin := r.Group(adminPath).Use(middleware.CheckIfAllowed())
	admin.POST("/new-device", h.adminHandler.CreateDevice)
	admin.DELETE("/delete-device/:deviceUUID", h.adminHandler.DeleteDevice)
	admin.PUT("/update-device", h.adminHandler.UpdateDevice)
	admin.PUT("/update-device-amount", h.adminHandler.UpdateDeviceAmount)

	// DEVICE ROUTES
	device := r.Group(devicesPath)
	device.GET("/", h.deviceHandler.GetAllDevices)
	device.GET("/title/:title", h.deviceHandler.GetDevicesByTitle)
	device.GET("/manufacturer/:manu", h.deviceHandler.GetDevicesByManufacturer)
	device.GET("/price", h.deviceHandler.GetDevicesByPrice)
	device.GET("/:uuid", h.deviceHandler.GetDeviceByUUID)
}

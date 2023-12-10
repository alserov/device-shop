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

func LoadRoutes(r *gin.Engine, h *Controller) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// AUTH
	userAuth := r.Group(authPath)
	userAuth.POST("/signup", h.authHandler.Signup)
	userAuth.POST("/login", h.authHandler.Login)

	// USER ACTIONS
	userActions := r.Group(actionsPath).Use(middleware.CheckIfAuthorized())
	userActions.PUT("/balance", h.userHandler.TopUpBalance)
	userActions.GET("/info/:userUUID", h.authHandler.GetInfo)
	userActions.POST("/new-favourite", h.collectionsHandler.AddToFavourite)
	userActions.DELETE("/delete-favourite", h.collectionsHandler.RemoveFromFavourite)
	userActions.GET("/favourite/:userUUID", h.collectionsHandler.GetFavourite)
	userActions.POST("/new-cart", h.collectionsHandler.AddToCart)
	userActions.DELETE("/delete-cart", h.collectionsHandler.RemoveFromCart)
	userActions.GET("/cart/:userUUID", h.collectionsHandler.GetCart)

	// ORDERS
	order := r.Group(ordersPath).Use(middleware.CheckIfAuthorized())
	order.POST("/new", h.orderHandler.CreateOrder)
	order.PUT("/update", h.orderHandler.UpdateOrder)
	order.GET("/:orderUUID", h.orderHandler.CheckOrder)

	// ADMIN routes
	admin := r.Group(adminPath).Use(middleware.CheckIfAllowed())
	admin.POST("/new-device", h.adminHandler.CreateDevice)
	admin.DELETE("/delete-device/:deviceUUID", h.adminHandler.DeleteDevice)
	admin.PUT("/update-device", h.adminHandler.UpdateDevice)
	admin.PUT("/update-device-amount", h.adminHandler.UpdateDeviceAmount)

	// DEVICE ROUTES
	device := r.Group(devicesPath)
	device.GET("/", h.devicesHandler.GetAllDevices)
	device.GET("/title/:title", h.devicesHandler.GetDevicesByTitle)
	device.GET("/manufacturer/:manu", h.devicesHandler.GetDevicesByManufacturer)
	device.GET("/price", h.devicesHandler.GetDevicesByPrice)
	device.GET("/:uuid", h.devicesHandler.GetDeviceByUUID)
}

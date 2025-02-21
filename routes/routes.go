package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/telexintegrations/Forex-Symbols-Notifier/handlers"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Deployed and Running fun nummbers app"})
	})

	
	router.GET("/integration.json", controller.Webhook)
	router.GET("/integration", controller.Webhook2)
	router.GET("/tick", controller.Get_symbols)
	router.POST("/tick", controller.Get_symbols)
}
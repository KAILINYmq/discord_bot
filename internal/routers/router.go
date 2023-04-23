package routers

import (
	"DiscordRolesBot/internal/routers/api"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.HandleMethodNotAllowed = true
	router.Use(gin.Recovery())

	v1 := router.Group("api/v1")
	v1.POST("bind/discord/:walletAddress", api.BindDiscord)

	return router
}

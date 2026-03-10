package routes

import (
	"vball/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	// PLAYER ROUTES
	player := router.Group("/admin")
	{
		player.GET("/", handlers.GetAllPlayersLoadOut_Handler)
		player.GET("/players/:steamid", handlers.GetPlayerBySteamID_Handler)
	}

	// ADMIN ROUTES (dashboard)
	admin := router.Group("/abilities")
	{
		admin.POST("/main", handlers.CreateMainAbility)
		admin.GET("/main", handlers.GetMainAbilities)
		admin.GET("/main/:id", handlers.GetMainAbility)
		admin.PATCH("/main/:id", handlers.UpdateMainAbility)
		admin.DELETE("/main/:id", handlers.DeleteMainAbility)

		admin.POST("/sub", handlers.CreateSubAbility)
		admin.GET("/sub", handlers.GetSubAbilities)
		admin.GET("/sub/:id", handlers.GetSubAbility)
		admin.PATCH("/sub/:id", handlers.UpdateSubAbility)
		admin.DELETE("/sub/:id", handlers.DeleteSubAbility)
	}

	// GAME RUNTIME ROUTES
	game := router.Group("/game")
	{
		/*
		   Request Body should be like this when calling /game/auth endpoint:
		   {
		     "steamId": "76561198000000001",
		     "username": "AceSpiker"
		   }

		   Endpoint Test
		   curl -X POST http://localhost:8080/game/auth -H "Content-Type: application/json" -d '{"steamId":"123","username":"Viku"}'
		*/
		game.POST("/auth", handlers.GetSteamLogin_Handler)

		/*
		   Endpoint Test
		   curl http://localhost:8080/game/abilities
		*/
		game.GET("/abilities", handlers.GetGameAbilities)
	}
}

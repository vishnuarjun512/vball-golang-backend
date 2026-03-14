package routes

import (
	"vball/internal/handlers"
	ability "vball/internal/handlers/abilities"
	vps "vball/internal/handlers/vps"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	// ADMIN ROUTES (dashboard)
	admin := router.Group("/abilities")
	{
		admin.POST("/main", ability.CreateMainAbility)
		admin.GET("/main", ability.GetMainAbilities)
		admin.GET("/main/:id", ability.GetMainAbility)
		admin.PATCH("/main/:id", ability.UpdateMainAbility)
		admin.DELETE("/main/:id", ability.DeleteMainAbility)

		admin.POST("/sub", ability.CreateSubAbility)
		admin.GET("/sub", ability.GetSubAbilities)
		admin.GET("/sub/:id", ability.GetSubAbility)
		admin.PATCH("/sub/:id", ability.UpdateSubAbility)
		admin.DELETE("/sub/:id", ability.DeleteSubAbility)
	}

	// PLAYER ROUTES
	player := router.Group("/admin")
	{
		player.GET("", handlers.GetAdminLoadOut_Handler)
		player.GET("/players/:steamid", handlers.GetPlayerBySteamID_Handler)
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
		/*
		   Endpoint Test
		   curl -X POST http://localhost:8080/game/matchmaking/join -H "Content-Type: application/json" -d '{"playerId":"steam_999","region":"asia"}'
		*/
		game.POST("/matchmaking/join", vps.JoinHandler)

		{
			/*
			   Endpoint Test
			   curl -X POST http://localhost:8080/game/matchmaking/leave -H "Content-Type: application/json" -d '{"playerId":"123"}'
			*/
		}
		game.POST("/matchmaking/leave", vps.LeaveHandler)

		/*
			LOGIC FOR SYNCING SERVER PLAYERS:
			1️⃣ delete all players for that server
			2️⃣ insert new player list
			3️⃣ update current_players

			Endpoint Test
			curl -X POST http://localhost:8080/game/server/sync -H "Content-Type: application/json" -d '{"serverId":4,"players":["steam_123","steam_456","steam_999"]}'
		*/

		game.POST("/server/sync", vps.SyncServer_Handler)
	}
}

package routes

import (
	"vball/internal/handlers"

	player_handler "vball/internal/tables/player"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	// PLAYER ROUTES
	player := router.Group("/admin")
	{
		player.GET("", handlers.GetAdminLoadOut_Handler)
		player.GET("/players/:steamid", player_handler.GetPlayerBySteamID_Handler)

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
		game.POST("/matchmaking/join", handlers.JoinHandler)

		{
			/*
			   Endpoint Test
			   curl -X POST http://localhost:8080/game/matchmaking/leave -H "Content-Type: application/json" -d '{"playerId":"123"}'
			*/
		}
		game.POST("/matchmaking/leave", handlers.LeaveHandler)

		/*
			LOGIC FOR SYNCING SERVER PLAYERS:
			1️⃣ delete all players for that server
			2️⃣ insert new player list
			3️⃣ update current_players

			Endpoint Test
			curl -X POST http://localhost:8080/game/server/sync -H "Content-Type: application/json" -d '{"serverId":4,"players":["steam_123","steam_456","steam_999"]}'
		*/

		game.POST("/server/sync", handlers.SyncServer_Handler)
	}
}

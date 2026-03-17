package routes

import (
	"vball/internal/handlers"

	"vball/internal/handlers/vps"

	mainAbility "vball/internal/tables/abilities/main"
	subAbility "vball/internal/tables/abilities/sub"
	"vball/internal/tables/gameserver"
	"vball/internal/tables/machine"
	"vball/internal/tables/region"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	// ADMIN ROUTES (dashboard)
	admin := router.Group("/abilities")
	{
		admin.POST("/main", mainAbility.CreateMainAbility)
		admin.GET("/main", mainAbility.GetMainAbilities)
		admin.GET("/main/:id", mainAbility.GetMainAbility)
		admin.PATCH("/main/:id", mainAbility.UpdateMainAbility)
		admin.DELETE("/main/:id", mainAbility.DeleteMainAbility)

		admin.POST("/sub", subAbility.UpdateSubAbility)
		admin.GET("/sub", subAbility.GetSubAbilities)
		admin.GET("/sub/:id", subAbility.GetSubAbility)
		admin.PATCH("/sub/:id", subAbility.UpdateSubAbility)
		admin.DELETE("/sub/:id", subAbility.DeleteSubAbility)
	}

	// PLAYER ROUTES
	player := router.Group("/admin")
	{
		player.GET("", handlers.GetAdminLoadOut_Handler)
		player.GET("/players/:steamid", handlers.GetPlayerBySteamID_Handler)

	}

	regions := router.Group("/region")
	{
		regions.GET("", region.GetAllRegions_Hanlder)
		regions.POST("", region.CreateRegion_Handler)
		regions.PATCH("/:id", region.UpdateRegion_Handler)
		regions.DELETE("/:id", region.DeleteRegion_Handler)
	}

	machines := router.Group("/machine")
	{
		machines.GET("", machine.GetAllMachines_Handler)
	}

	game_server := router.Group(("/gameserver"))
	{
		game_server.POST("", gameserver.CreateGameServer_Handler)
		game_server.GET("/:id", gameserver.GetGameServer_Handler)
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

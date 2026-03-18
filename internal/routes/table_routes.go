package routes

import (
	mainAbility "vball/internal/tables/abilities/main"
	subAbility "vball/internal/tables/abilities/sub"
	"vball/internal/tables/gameserver"
	"vball/internal/tables/machine"
	"vball/internal/tables/player"
	"vball/internal/tables/region"

	"github.com/gin-gonic/gin"
)

func SetupTableRoutes(router *gin.Engine) {

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

	regions := router.Group("/region")
	{
		regions.GET("", region.GetAllRegions_Hanlder)
		regions.POST("", region.CreateRegion_Handler)
		regions.GET("/:id", region.GetRegion_Handler)
		regions.PATCH("/:id", region.UpdateRegion_Handler)
		regions.DELETE("/:id", region.DeleteRegion_Handler)
	}

	machine_route := router.Group("/machine")
	{
		machine_route.GET("", machine.GetAllMachines_Handler)
		// machine_router.POST("", region.CreateMachine_Handler)
		// machine_router.GET("/:id", region.GetMachine_Handler)
		// machine_router.PATCH("/:id", region.UpdateMachine_Handler)
		// machine_router.DELETE("/:id", region.DeleteMachine_Handler)
	}

	game_server := router.Group(("/gameserver"))
	{
		game_server.POST("", gameserver.CreateGameServer_Handler)
		game_server.GET("/:id", gameserver.GetGameServer_Handler)
	}

	player_route := router.Group("/player")
	{
		player_route.POST("", player.CreatePlayer_Handler)
		player_route.DELETE("/:id", player.DeletePlayer_Handler)
	}

}

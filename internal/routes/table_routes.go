package routes

import (
	mainAbility "vball/internal/tables/abilities/main"
	subAbility "vball/internal/tables/abilities/sub"
	"vball/internal/tables/gameserver"
	"vball/internal/tables/machine"
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

}

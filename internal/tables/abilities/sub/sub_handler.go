package subAbility

import (
	"strconv"

	"vball/utils"

	"github.com/gin-gonic/gin"
)

func CreateSubAbility(c *gin.Context) {
	var ability SubAbility
	if err := c.ShouldBindJSON(&ability); err != nil {
		utils.SendError(c, 400, "Invalid request", err)
		return
	}

	if err := CreateSubAbility_Service(ability); err != nil {
		utils.SendError(c, 500, "Creation failed", err)
		return
	}
	c.JSON(201, gin.H{"error": false, "message": "Sub ability created"})
}

func GetSubAbilities(c *gin.Context) {
	abilities, err := GetSubAbilities_Service()
	if err != nil {
		utils.SendError(c, 500, "Failed to fetch abilities", err)
		return
	}
	c.JSON(200, gin.H{"error": false, "abilities": abilities})
}

func GetSubAbility(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	ability, err := GetSubAbility_Service(id)
	if err != nil {
		utils.SendError(c, 404, "Sub ability not found", err)
		return
	}
	c.JSON(200, gin.H{"error": false, "ability": ability})
}

func UpdateSubAbility(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var ability SubAbility
	if err := c.ShouldBindJSON(&ability); err != nil {
		utils.SendError(c, 400, "Invalid input", err)
		return
	}

	if err := UpdateSubAbility_Service(id, ability); err != nil {
		utils.SendError(c, 500, "Update failed", err)
		return
	}
	c.JSON(200, gin.H{"error": false, "message": "Updated Sub Ability"})
}

func DeleteSubAbility(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteSubAbility_Service(id); err != nil {
		utils.SendError(c, 500, "Delete failed", err)
		return
	}
	c.JSON(200, gin.H{"error": false, "message": "Deleted Sub Ability"})
}

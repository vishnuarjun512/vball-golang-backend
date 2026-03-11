package handlers

import (
	"strconv"

	"vball/internal/models"
	"vball/internal/services"

	"github.com/gin-gonic/gin"
)

func CreateMainAbility(c *gin.Context) {

	var ability models.MainAbility

	if err := c.BindJSON(&ability); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	err := services.CreateMainAbility(ability)

	if err != nil {
		c.JSON(500, gin.H{"error": "creation failed"})
		return
	}

	c.JSON(200, gin.H{"message": "ability created"})
}

func GetMainAbilities(c *gin.Context) {

	abilities, err := services.GetMainAbilities()

	if err != nil {
		c.JSON(500, gin.H{"error": "failed"})
		return
	}

	c.JSON(200, abilities)
}

func GetMainAbility(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	ability, err := services.GetMainAbility(id)

	if err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}

	c.JSON(200, ability)
}

func UpdateMainAbility(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	var ability models.MainAbility

	c.BindJSON(&ability)

	err := services.UpdateMainAbility(id, ability)

	if err != nil {
		c.JSON(500, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, gin.H{"message": "updated"})
}

func DeleteMainAbility(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	err := services.DeleteMainAbility(id)

	if err != nil {
		c.JSON(500, gin.H{"error": "delete failed"})
		return
	}

	c.JSON(200, gin.H{"message": "deleted"})
}

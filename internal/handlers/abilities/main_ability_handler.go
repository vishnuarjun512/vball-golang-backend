package handlers

import (
	"fmt"
	"strconv"

	"vball/internal/models"
	"vball/internal/services"

	"github.com/gin-gonic/gin"
)

func CreateMainAbility(c *gin.Context) {

	var newAbility models.CreateAbilityRequest

	if err := c.ShouldBindJSON(&newAbility); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	createdAbility, err := services.CreateMainAbility(newAbility)

	if err != nil {
		fmt.Println("Error creating main ability:", err)
		c.JSON(500, gin.H{"error": "creation failed"})
		return
	}

	fmt.Println("Main ability created successfully")
	c.JSON(200, gin.H{"message": "ability created", "ability": createdAbility})
}

func GetMainAbilities(c *gin.Context) {

	abilities, err := services.GetMainAbilities()

	if err != nil {
		fmt.Println("Error fetching main abilities:", err)
		c.JSON(500, gin.H{"error": "failed"})
		return
	}

	fmt.Println("Main abilities fetched successfully")
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
		fmt.Println("Error updating main ability:", err)
		c.JSON(500, gin.H{"message": "update failed", "error": true})
		return
	}

	fmt.Println("Main ability updated successfully")
	c.JSON(200, gin.H{"message": "updated", "error": false})
}

func DeleteMainAbility(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	err := services.DeleteMainAbility(id)

	if err != nil {
		fmt.Println("Error deleting main ability:", err)
		c.JSON(500, gin.H{"error": true, "message": "deletion failed"})
		return
	}

	fmt.Println("Main ability deleted successfully")
	c.JSON(200, gin.H{"message": "deleted", "error": false})
}

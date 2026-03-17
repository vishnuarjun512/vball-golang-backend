package subAbility

import (
	"strconv"
	"vball/internal/models"

	"github.com/gin-gonic/gin"
)

func CreateSubAbility(c *gin.Context) {

	var ability models.SubAbility

	if err := c.BindJSON(&ability); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	err := CreateSubAbility_Service(ability)

	if err != nil {
		c.JSON(500, gin.H{"error": "creation failed"})
		return
	}

	c.JSON(200, gin.H{"message": "sub ability created"})
}

func GetSubAbilities(c *gin.Context) {

	abilities, err := GetSubAbilities_Service()

	if err != nil {
		c.JSON(500, gin.H{"error": "failed"})
		return
	}

	c.JSON(200, abilities)
}

func GetSubAbility(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	ability, err := GetSubAbility_Service(id)

	if err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}

	c.JSON(200, ability)
}

func UpdateSubAbility(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	var ability models.SubAbility

	c.BindJSON(&ability)

	err := UpdateSubAbility_Service(id, ability)

	if err != nil {
		c.JSON(500, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, gin.H{"message": "updated"})
}

func DeleteSubAbility(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	err := DeleteSubAbility_Service(id)

	if err != nil {
		c.JSON(500, gin.H{"error": "delete failed"})
		return
	}

	c.JSON(200, gin.H{"message": "deleted"})
}

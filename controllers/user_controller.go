package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/maanxester/webapi-golang/database"
	"github.com/maanxester/webapi-golang/models"
	"strconv"
)

func ShowUsers(c *gin.Context) {
	id := c.Param("id") // Retorna uma string

	newid, err := strconv.Atoi(id) // Converte para int

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	var user models.User

	err = db.First(user, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find user:" + err.Error() ,
		})
		return
	}

	c.JSON(200, user)
}

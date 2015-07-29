package controllers

import (
	"github.com/ObamaPhony/rest-api/models"
	"github.com/gin-gonic/gin"
)

func SGenListSpeeches(c *gin.Context) {
	speechesList := models.SpeechesList{
		ID:          1,
		SpeechTitle: "I'm a president yay!",
		President:   "Barack Obama",
	}

	c.JSON(200, speechesList)
}

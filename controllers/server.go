package controllers

import (
	// "github.com/ObamaPhony/obama-rest-api/models"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"time"
)

func StartServer(hostname string) error {
	r := gin.Default()

	r.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))

	v1 := r.Group("api/v1")
	{
		v1.GET("/speechgen/speeches", SGenListSpeeches)
	}

	err := r.Run(hostname)
	if err != nil {
		return err
	}

	return nil
}

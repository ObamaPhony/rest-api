package controllers

import (
	// "github.com/ObamaPhony/rest-api/models"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

func StartServer(hostname string, wg *sync.WaitGroup, chanErrorResult chan error) {
	r := gin.Default() // Create a instance of Gin.

	r.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true)) // Use the Ginrus middleware logger.

	v1 := r.Group("api/v1")
	{
		v1.GET("/speechgen/speeches", SGenListSpeeches)
	}

	err := r.Run(hostname)
	if err != nil {
		chanErrorResult <- err
		wg.Done()
	}

	chanErrorResult <- nil
	wg.Done()
}

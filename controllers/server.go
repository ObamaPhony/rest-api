package controllers

import (
	// "github.com/ObamaPhony/rest-api/models"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

func StartServer(hostname string, quit chan bool, wg *sync.WaitGroup) error {
	done := make(chan bool) // Create a done channel of bool type to return when the Server is finished.

	r := gin.Default() // Create a instance of Gin

	r.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))

	v1 := r.Group("api/v1")
	{
		v1.GET("/speechgen/speeches", SGenListSpeeches)
	}

	err := r.Run(hostname)
	if err != nil {
		done <- true
		wg.Done()
		return err
	}

	done <- true
	wg.Done()

	return nil
}

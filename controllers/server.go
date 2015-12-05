package controllers

import (
	"errors"
	"github.com/ObamaPhony/rest-api/models"
	"github.com/pilu/traffic"
	"net/http"
	"sync"
)

// StartServer starts the REST API. It takes three arguments, a hostname of
// string type (This should be formatted like '127.0.0.1:8080', or else the
// system will panic!), a waitgroup from the sync package (this is sent from
// the server.go file)
func StartServer(listen string, wg *sync.WaitGroup, chanErrorResult chan error) {
	loggers := models.ReturnLoggers()

	// Diagnostics, hostname len == 0?
	if len(listen) == 0 {
		err := errors.New("HOSTNAME_DIAGCHECK_LEN_EQ0")
		loggers.LogControllers.Error("The REST API has been passed a empty hostname string! We can't continue, check the config.json file", "The error message in all it's error type glory", err)
		panic(err)
	}

	loggers.LogControllers.Debug("Creating a router instance..")

	// Create router instance.
	router := traffic.New()

	router.Get("/speeches", GenSpeech)
	router.Post("/speech", GenSpeech)

	loggers.LogControllers.Info("Starting the REST API...")
	err := http.ListenAndServe(listen, router)
	if err != nil {
		loggers.LogControllers.Error("The REST API has encountered a error during listening time.", "The errro rmessage in all it#s error type glory", err)

		chanErrorResult <- err
		wg.Done()
	}

	// Generally the server is stopped with CTRL-C, but I would like to
	// implement a kill signal at some point that the server detects
	// automatically. (TODO)

	wg.Done()
}

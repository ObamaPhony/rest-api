package main

import (
	"bytes"
	"fmt"
	"github.com/ObamaPhony/rest-api/config"
	"github.com/ObamaPhony/rest-api/controllers"
	"github.com/ObamaPhony/rest-api/exec"
	"github.com/ObamaPhony/rest-api/models"
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"sync"
)

// Global Configuration file.
var Config *viper.Viper

// Loggers
var LogBase *log.Entry
var LogConfig *log.Entry
var LogControllers *log.Entry
var LogExec *log.Entry
var LogSpeech *log.Entry

func init() {
	var err error

	// Fill in the logger instances into the variables previously definewd.
	LogBase = models.ReturnBaseLogger()
	LogConfig = models.ReturnLogConfig()
	LogControllers = models.ReturnLogControllers()
	LogExec = models.ReturnLogExec()
	LogSpeech = models.ReturnLogSpeech()

	// Define the environment variables.
	orestCONFIGPATH := os.Getenv("orestCONFIGPATH")
	orestCONFIGNAME := os.Getenv("orestCONFIGNAME")

	// Test if the ENV variables are empty, if so spontaneously die.

	if len(orestCONFIGNAME) == 0 {
		fmt.Printf("The orestCONFIGNAME environment variable is NOT found. The REST API cannot continue, *shoots in the head*\n") // Possibly not PG at all.
		panic("No ENV.")
	}

	if len(orestCONFIGPATH) == 0 {
		fmt.Printf("The orestCONFIGPATH environment variable is NOT found. The REST API cannot continue, *shoots in the head*\n") // Possibly not PG at all.
		panic("No ENV.")
	}

	Config, err = config.GetViper(orestCONFIGPATH, orestCONFIGNAME)
	if err != nil {
		fmt.Printf("Error found when initalizing the configuration instance. Error message: %s\n. Cannot continue, *shoots in the head*"+" Error message: %s\n", err)
	}

	LogConfig.Info("Loaded the configuration instance.")
	LogBase.Info("Seems OK, deploying REST server.")

}

func main() {
	/** I want to separate the following into a subpackage at some point! **/
	/** For now keeping it here for reference and to showcase my async. **/
	var wg sync.WaitGroup

	chanSpeechBytesBuffer := make(chan *bytes.Buffer) // Create a channel for the bytes Buffer from the Python natural language processor/function.
	chanSpeechError := make(chan error)               // Create a channel for errors returned from the Python natural language processor/function.

	chanRESTError := make(chan error) // Create a channel for errors returned from the REST server.

	wg.Add(2)

	go exec.SpeechAnalysis(chanSpeechBytesBuffer, chanSpeechError, false, &wg, "", "")
	go controllers.StartServer(":8080", &wg, chanRESTError)

	// SpeechBytesBufferResult := <-chanSpeechBytesBuffer // Don't need this.. for now.
	SpeechErrorResult := <-chanSpeechError
	if SpeechErrorResult != nil {
		panic(SpeechErrorResult)
	}

	RESTErrorResult := <-chanRESTError
	if RESTErrorResult != nil {
		panic(RESTErrorResult)
	}

	wg.Wait()
}

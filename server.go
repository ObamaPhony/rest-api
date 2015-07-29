package main

import (
	"bytes"
	"fmt"
	"github.com/ObamaPhony/rest-api/config"
	"github.com/ObamaPhony/rest-api/controllers"
	"github.com/ObamaPhony/rest-api/exec"
	"github.com/spf13/viper"
	"os"
	"sync"
)

var Config *viper.Viper

func init() {
	var err error

	OREST_CONFIGPATH := os.Getenv("OREST_CONFIGPATH")
	OREST_CONFIGNAME := os.Getenv("OREST_CONFIGNAME")
	// Test if the ENV variables are empty, if so spontaneously die.

	if len(OREST_CONFIGNAME) == 0 {
		fmt.Printf("The OREST_CONFIGNAME environment variable is NOT found. The REST API cannot continue, *shoots in the head*\n") // Possibly not PG at all.
		panic("No ENV.")
	}

	if len(OREST_CONFIGPATH) == 0 {
		fmt.Printf("The OREST_CONFIGPATH environment variable is NOT found. The REST API cannot continue, *shoots in the head*\n") // Possibly not PG at all.
		panic("No ENV.")
	}

	Config, err = config.GetViper(os.Getenv("OREST_CONFIGPATH"), os.Getenv("OREST_CONFIGNAME"))
	if err != nil {
		fmt.Printf("Error found when initalizing the configuration instance. Error message: %s\n. Cannot continue, *shoots in the head*"+" Error message: %s\n", err)
	}
}

func main() {
	/** I want to seperate the following into a subpackage at some point! **/
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

	fmt.Println("Finished?")
}

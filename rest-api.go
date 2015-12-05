package main

import (
	"flag"
	"fmt"
	"github.com/ObamaPhony/rest-api/config"
	"github.com/ObamaPhony/rest-api/controllers"
	"github.com/ObamaPhony/rest-api/models"
	"github.com/jeffail/gabs"
	"io/ioutil"
	"os"
	"sync"
)

// Config is the local configuration instance of a gabs container.
var Config *gabs.Container

// Loggers is the local instance of the 'Loggers' struct from the models package.
var Loggers models.Loggers


var (
	// ConfigFileFlag is the value of the 'ConfigFile' flag passed (optionally, instead of the ENV variable to the API)
	ConfigFileFlag = flag.String("ConfigFile", "", "This flag is the location of the configuration file for the REST API")
)

func init() {
	// Get Loggers struct
	Loggers = models.ReturnLoggers()

	// Parse flags
	flag.Parse()

	// Define the environment variables.
	ConfigFile := os.Getenv("OBAMA_RESTAPI_CONFIG")

	if len(ConfigFile) == 0 {

		Loggers.LogConfig.Debug("The OBAMAP_RESTAPI_CONFIG environment variable has not been filled.")
		Loggers.LogConfig.Debug("Instead, trying to access the Config file location from cmd flags.")

		// ConfigFile should now have the value of ConfigFileFlag, testing again.
		ConfigFile = *ConfigFileFlag

		if len(ConfigFile) == 0 {
			Loggers.LogConfig.Fatal("The OBAMA_RESTAPI_CONFIG environment variable or command-line flag has still not been found. Having to stop execution of API!")
		}
	}

	content, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		fmt.Printf("Error found when initalizing the configuration instance. Error message: %s\n. Cannot continue.\n", err)
		panic("No CONFIG AVAILABLE")
	}

	Config, err = config.ReturnGABS(content)
	if err != nil {
		fmt.Printf("Error found when initalizing the configuration instance. Error message: %s\n. Cannot continue.\n", err)
		panic("No CONFIG AVAILABLE")
	}

	version := Config.Path("version").Data().(string)

	Loggers.LogConfig.Info("Loaded the configuration.", "ObamaPhony REST-API Version", version)
}

func main() {
	listen := Config.Path("server.listenAddr").Data().(string) +
		Config.Path("server.listenPort").Data().(string)
	var wg sync.WaitGroup // Create a waitgroup instance to manage the goroutines.

	chanRESTError := make(chan error) // Create a channel for errors returned from the REST server.

	wg.Add(1) // Add one goroutine, that's all we have at the moment.

	go controllers.StartServer(listen, &wg, chanRESTError) // Create a goroutine that runs the REST server in the background, and pass the port, channel for error
	// handling and a reference pointer to the waitgroup we created before.

	RESTErrorResult := <-chanRESTError // Transfer the REST Error channel into a variable.
	if RESTErrorResult != nil {
		panic(RESTErrorResult)
	}

	wg.Wait()
}

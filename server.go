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
	var w sync.WaitGroup

	chanbufferspeech := make(chan *bytes.Buffer, 1)
	chanerrorspeech := make(chan error, 1)
	chandonespeech := make(chan bool, 1)

	chandonerest := make(chan bool, 1)

	w.Add(2)

	go exec.SpeechAnalysis(chanbufferspeech, chanerrorspeech, chandonespeech, true, &w, "", "")
	go controllers.StartServer(":8080", chandonerest, &w)

	chanbufferspeechResult := <-chanbufferspeech
	chanerrorspeechResult := <-chanerrorspeech
	chandonespeechResult := <-chandonespeech
	chandonerestResult := <-chandonerest

	if chanerrorspeechResult != nil {
		panic(chanerrorspeechResult)
	}

	if chandonerestResult == false {
		fmt.Println("rest not done yet.")
	}

	if chandonespeechResult == false {
		fmt.Println("speech not done yet.")
	}

	fmt.Println(chanbufferspeechResult.String())
	w.Wait()
}

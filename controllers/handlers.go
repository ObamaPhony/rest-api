package controllers

import (
	"github.com/ObamaPhony/rest-api/exec"
	"github.com/ObamaPhony/rest-api/models"
	"github.com/jeffail/gabs"
	"github.com/pilu/traffic"
	"io/ioutil"
	"os"
)

// GenSpeech does
func GenSpeech(w traffic.ResponseWriter, r *traffic.Request) {
	loggers := models.ReturnLoggers()

	// Create channels.
	loggers.LogControllers.Debug("Channels being generated.")
	loggers.LogControllers.Debug("Channels generated.")

	SpeechRequest := r.Body

	// Create config instance
	content, err := ioutil.ReadFile(os.Getenv("OBAMAP_RESTAPI_CONFIG"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.WriteJSON(models.ErrorJSON{
			500,
			err,
		})
	}
	config, err := gabs.ParseJSON(content)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.WriteJSON(models.ErrorJSON{
			500,
			err,
		})
	}
	FileOutputPath := config.Path("speech.fileOutputPath").Data().(string)
	PythonLocation := config.Path("speech.pythonLocation").Data().(string)
	FileOutput := config.Path("speech.fileOutput").Data().(bool)

	loggers.LogControllers.Info("Executing Speech Analysis.")

	//exec.SpeechAnalysis(&a, done)
	result, err := exec.SpeechAnalysis2(exec.Arguments{
		FileOutputPath: FileOutputPath,
		ScriptLocation: PythonLocation,
		FileOutput:     FileOutput,
		SpeechRequest:  SpeechRequest,
	})

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.WriteJSON(models.ErrorJSON{
			500,
			err,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteText(result)
}

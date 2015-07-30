package models

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jeffail/gabs"
)

var LogObamaREST = log.WithFields(log.Fields{"App": "ObamaPhony/REST"})

var (
	LogConfig      = LogObamaREST.WithFields(log.Fields{"Module": "/Config"})
	LogControllers = LogObamaREST.WithFields(log.Fields{"Module": "/Controllers"})
	LogExec        = LogObamaREST.WithFields(log.Fields{"Module": "/Exec"})
	LogSpeech      = LogObamaREST.WithFields(log.Fields{"Module": "/Exec/Speech"})
)

type SpeechesList struct {
	ID       int        `json:"id"`
	name     string     `json:"name"`
	analysis [][]string `json:"analysis"`
}

func ReturnBaseLogger() *log.Entry {
	return LogObamaREST
}

func ReturnLogConfig() *log.Entry {
	return LogConfig
}

func ReturnLogControllers() *log.Entry {
	return LogControllers
}

func ReturnLogExec() *log.Entry {
	return LogExec
}

func ReturnLogSpeech() *log.Entry {
	return LogSpeech
}

func GABSParseJSON(json []byte) (*gabs.Container, error) {
	jsonParsed, err := gabs.ParseJSON(json)
	if err != nil {
		return jsonParsed, err
	}

	// No errors, we're good here.
	return jsonParsed, nil
}

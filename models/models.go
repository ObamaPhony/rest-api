package models

import (
	log "github.com/Sirupsen/logrus"
)

var LogObamaREST = log.WithFields(log.Fields{"App": "ObamaPhony/REST"})

type SpeechesList struct {
	ID          int    `json:"id"`
	SpeechTitle string `json:"speechTitle"`
	President   string `json:"president"`
}

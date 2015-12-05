package models

import (
	log "github.com/mgutz/logxi/v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// SpeechJSON is a struct defining the structure of each speech
type SpeechJSON struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Title    string        `json:"title" bson:"title"`
	Analysis string        `json:"analysis" bson:"analysis"`
}

// ErrorJSON d
type ErrorJSON struct {
	Status int   `json:"status"`
	Error  error `json:"error"`
}

// Loggers is a struct defining the structure of loggers.
type Loggers struct {
	LogBase        log.Logger
	LogConfig      log.Logger
	LogControllers log.Logger
	LogExec        log.Logger
	LogDB          log.Logger
	LogModels      log.Logger
}

// ReturnLoggers returns the Loggers struct with initalized loggers.
func ReturnLoggers() Loggers {
	loggers := Loggers{
		LogBase:        log.New("REST/Base"),
		LogConfig:      log.New("REST/Config"),
		LogControllers: log.New("REST/Controllers"),
		LogExec:        log.New("REST/Exec"),
		LogDB:          log.New("REST/DB"),
		LogModels:      log.New("REST/Models"),
	}

	return loggers
}

// GetMongo returns a MongoDB session after connecting to the database.
func GetMongo(hostname string, monotonic bool) (*mgo.Session, error) {
	session, err := mgo.Dial(hostname)
	if err != nil {
		return session, err
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, monotonic)

	return session, nil
}

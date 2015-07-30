package db

import (
	// "github.com/ObamaPhony/rest-api/models"
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)

func GetMongo(hostname string, monotonic bool) (error, *mgo.Session) {
	session, err := mgo.Dial(hostname)
	if err != nil {
		return err, session
	}

	defer session.Close()
	return nil, session
}

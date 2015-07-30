package db

import (
	"github.com/ObamaPhony/rest-api/models"
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)

func GetMongo(hostname string, monotonic bool) (error, *mgo.Session) {
	session, err := mgo.Dial(hostname)
	if err != nil {
		return err, session
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, monotonic)

	return nil, session
}

func AddDocumentToMongo(session *mgo.Session, database string, collection string, speechListModel *models.SpeechesList) error {
	c := session.DB(database).C(collection)
	err := c.Insert(speechListModel)
	if err != nil {
		return err
	}

	return nil
}

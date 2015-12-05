package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func TestSpec(t *testing.T) {
	Convey("Module: Models Testing", t, func() {
		var loggers Loggers

		Convey("Testing the Loggers return function", func() {

			loggers = ReturnLoggers()

			Convey("The loggers struct returned should be of type 'models.Loggers'", func() {
				So(reflect.TypeOf(loggers).String(), ShouldEqual, "models.Loggers")
			})

		})

		Convey("Testing the MongoDB session return function", func() {
			session, err := GetMongo("localhost", false)

			Convey("The MongoDB session returned should be of type '*mgo.Session'", func() {
				So(reflect.TypeOf(session).String(), ShouldEqual, "*mgo.Session")
			})

			Convey("The MongoDB session function should not have returned a error", func() {
				So(err, ShouldEqual, nil)
			})
		})

	})
}

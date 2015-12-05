package config

import (
	"github.com/jeffail/gabs"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestSpec(t *testing.T) {
	Convey("Module: config", t, func() {
		var Config *gabs.Container

		Convey("When requesting a GABS Configuration instance", func() {
			// Load config.json file into []bytes (content).
			content, err := ioutil.ReadFile(os.Getenv("PWD") + "/config.json")
			if err != nil {
				Convey("The error returned should be nil", func() {
					So(err, ShouldEqual, nil)
				})
			}

			Config, err = ReturnGABS(content)
			if err != nil {
				Convey("The error returned should be nil", func() {
					So(err, ShouldEqual, nil)
				})
			}

			Convey("The Config instance should be of type '*gabs.Container'", func() {
				So(reflect.TypeOf(Config).String(), ShouldEqual, "*gabs.Container")

			})

		})

	})
}

package config

import (
	"github.com/jeffail/gabs"
)

// ReturnGABS returns a gabs container for configuration parsing.
func ReturnGABS(file []byte) (*gabs.Container, error) {
	jsonParsed, err := gabs.ParseJSON(file)
	if err != nil {
		return jsonParsed, err
	}

	return jsonParsed, nil
}

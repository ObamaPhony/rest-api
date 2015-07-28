package exec

import (
	"bytes"

	"gopkg.in/pipe.v2"
)

func SpeechAnalysisWriteFileStdout() (*bytes.Buffer, error, chan *bytes.Buffer, chan error) {
	// Define the byes buffer
	b := &bytes.Buffer{}
	p := pipe.Line(
		pipe.Exec("python ../../speech-analysis/analyse-speeches.py"),
		pipe.Tee(b),
		pipe.WriteFile("../result.json", 0644),
	)
	err := pipe.Run(p)
	// Channels
	cb1 := make(chan *bytes.Buffer, 1)
	cb2 := make(chan error, 1)
	if err != nil {
		return b, err, cb1, cb2
	}

	return b, nil, cb1, cb2
}

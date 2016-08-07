package exec

import (
	//	log "github.com/inconshreveable/log15"
	"bytes"
	"encoding/json"
	"gopkg.in/pipe.v2"
	"io"
	"os"
	"time"
)

type SAArguments struct {
	FileOUTPATH string
	SAScriptLOC string
	FileOUT     bool
	SpeechREQ   io.ReadCloser
}

// SAReturnASYS takes the SAArguments struct, and returns the analysis from the speech analysis program
func SAReturnASYS(a *SAArguments) (errresult error, result string) {

	if a.FileOUT == true {
		// We're outputting this to a file.
		// Create new bytes.Buffer instance
		BUFFER := new(bytes.Buffer)
		// Define the filename to output to.
		FILENAME := a.FileOUTPATH + "/" + "speechoutput_" + time.Now().Format(time.RFC3339) + ".json"

		// Create the output file.
		file, err := os.Create(FILENAME)
		if err != nil {
			return err, ""
		}
		defer file.Close()

		// Prepare the speech-analysis script pipe.
		p := pipe.Line(
			pipe.Read(a.SpeechREQ),
			pipe.Exec(a.SAScriptLOC),
			pipe.Tee(BUFFER),
		)

		err = pipe.Run(p)

		if err != nil {
			return err, ""
		}

		outputBytes := &bytes.Buffer{}
		if err := json.Compact(outputBytes, BUFFER.Bytes()); err != nil {
			panic(err)
		}

		return nil, outputBytes.String()
	}

	return nil, ""

}

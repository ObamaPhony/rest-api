package exec

import (
	"bytes"
	"gopkg.in/pipe.v2"
	"sync"
)

// SpeechAnalysis does takes three channels, one to return the buffer, one to return any errors and another to let the calling function know when it's done.
// It also takes a bool variable to signify if we are writing to file from the output of the python script as well as to a variable.
func SpeechAnalysis(chanBytesBuffer chan *bytes.Buffer, chanErrorResult chan error, file bool, wg *sync.WaitGroup, speechlocationFile string, speechlocationVar string) {

	switch file {
	// A true
	case true:
		// Initalize the buffer struct.
		b := &bytes.Buffer{}

		// Create a Pipe instance to pipe to the speech-analysis module.
		p := pipe.Line(
			pipe.ReadFile("./test-speeches.txt"), // TODO: Probs should look at functions arguments for this.
			// pipe.Exec("python", "/home/dzrodrig/ObamaPhony/speech-analysis/analyse-speeches.py"), // The full path should be changed in production!
			pipe.Exec("python", "/home/dzrodrig/dev/ObamaPhony/speech-analysis/analyse-speeches.py"),
			pipe.Tee(b),                          // Output to the 'b' buffer.
			pipe.WriteFile("results.json", 0644), // Also output to the file, we want that to happen!
		)
		// Run the Pipe instance.
		err := pipe.Run(p)
		// Error handling for the win.
		if err != nil {
			// If error detected, send back two channels previously defined, one with the buffer result (most likely corrupted or failed, but let's do it anyway)
			// and a channel with the error value.
			chanBytesBuffer <- b
			chanErrorResult <- err
			wg.Done()
		} else {
			// If no error detected, send back buffer instance and nil error value.
			chanBytesBuffer <- b
			chanErrorResult <- nil
			wg.Done()
		}
	case false:
		// If the file bool is false, we shouldn't write to a file, and just to the buffer and return it.
		b := &bytes.Buffer{}
		p := pipe.Line(
			pipe.ReadFile("./test-speeches.txt"), // TODO: Probs should look at functions arguments for this.
			// pipe.Exec("python", "/data1/_NASDrive/dzrodrig/Documents/dev/projects/src/github.com/ObamaPhony/speech-analysis/analyse-speeches.py"), // The full path should be changed in production!
			pipe.Exec("python", "/home/dzrodrig/dev/ObamaPhony/speech-analysis/analyse-speeches.py"),
			pipe.Tee(b), // Outputting to the 'b' buffer.
		)

		err := pipe.Run(p)
		if err != nil {
			// If error detected, send back two channels previously defined, one with the buffer result (most likely corrupted or failed, but let's do it anyway)
			// and a channel with the error value.
			chanBytesBuffer <- b
			chanErrorResult <- err
			wg.Done()
		} else {
			// If no error detected, send back buffer instance and nil error value.
			chanBytesBuffer <- b
			chanBytesBuffer <- nil
			wg.Done()
		}
	}
}

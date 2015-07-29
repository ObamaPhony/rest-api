package exec

import (
	"bytes"
	"gopkg.in/pipe.v2"
)

func SpeechAnalysis(cb1 chan *bytes.Buffer, ce1 chan error, file bool) {
	cb1 = make(chan *bytes.Buffer, 1)
	ce1 = make(chan error, 1)

	println("starting")
	if file == true {
		println("true")
		// Initalize the buffer struct.
		b := &bytes.Buffer{}

		// Create a Pipe instance to pipe to the speech-analysis module.
		p := pipe.Line(
			pipe.ReadFile("./test-speeches.txt"),                                                                                                  // TODO: Probs should look at functions arguments for this.
			pipe.Exec("python", "/data1/_NASDrive/dzrodrig/Documents/dev/projects/src/github.com/ObamaPhony/speech-analysis/analyse-speeches.py"), // The full path should be changed in production!
			pipe.Tee(b),                          // Output to the 'b' buffer.
			pipe.WriteFile("results.json", 0644), // Also output to the file, we want that to happen!
		)

		// Run the Pipe instance.
		err := pipe.Run(p)
		// Error handling for the win.
		if err != nil {
			// If error detected, send back two channels previously defined, one with the buffer result (most likely corrupted or failed, but let's do it anyway)
			// and a channel with the error value.
			cb1 <- b
			ce1 <- err
		} else {
			// If no error detected, send back buffer instance and nil error value.
			cb1 <- b
			ce1 <- nil
			println("written.")
		}

	} else if file == false {
		println("false")
		// If the file bool is false, we shouldn't write to a file, and just to the buffer and return it.
		b := &bytes.Buffer{}
		p := pipe.Line(
			pipe.ReadFile("./test-speeches.txt"),                                                                                                  // TODO: Probs should look at functions arguments for this.
			pipe.Exec("python", "/data1/_NASDrive/dzrodrig/Documents/dev/projects/src/github.com/ObamaPhony/speech-analysis/analyse-speeches.py"), // The full path should be changed in production!
			pipe.Tee(b), // Outputting to the 'b' buffer.
		)

		err := pipe.Run(p)
		if err != nil {
			// If error detected, send back two channels previously defined, one with the buffer result (most likely corrupted or failed, but let's do it anyway)
			// and a channel with the error value.
			cb1 <- b
			ce1 <- err
			return
		} else {
			// If no error detected, send back buffer instance and nil error value.
			cb1 <- b
			ce1 <- nil
			println("written")
			return
		}

	}

}

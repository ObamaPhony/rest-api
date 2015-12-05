package exec

import (
	"bytes"
	"encoding/json"
	"gopkg.in/pipe.v2"
	"io"
)

// Arguments is a struct
type Arguments struct {
	FileOutputPath string
	ScriptLocation string
	FileOutput     bool
	SpeechRequest  io.ReadCloser
}

/*// SpeechAnalysis does takes three channels, one to return the buffer, one to return any errors and another to let the calling function know when it's done.*/
//// It also takes a bool variable to signify if we are writing to file from the output of the python script as well as to a variable.
//func SpeechAnalysis(a *Arguments, done chan bool) {
//// If FileOutput is true, start a new goroutine and continue the analysis, and
//// then end the function.
//if a.FileOutput == true {
//// Create a new buffer
//buffer := new(bytes.Buffer)
//fileName := a.FileOutputPath + "/" + "speechOutput_" + time.Now().Format(time.RFC3339) + ".json"

//// Create the output file, as specified.
//file, err := os.Create(fileName)
//if err != nil {
//a.ChanError <- err
//close(done)
//}
//defer file.Close()

//// Prepare the speech analysis pipe.
//p := pipe.Line(
//pipe.ReadFile("/home/dzrodrig/dev/projects/src/github.com/ObamaPhony/speech-analysis/test-speeches.txt"),
//pipe.Exec("python3", a.ScriptLocation),
//)
//// Execute the pipe.
//output, err := pipe.CombinedOutput(p)
//if err != nil {
//a.ChanError <- err
//close(done)
//}

//// Convert the pipe output into a buffer.
//if err := json.Compact(buffer, output); err != nil {
//a.ChanError <- err
//close(done)
//}

//// Pass the buffer to the channel
//a.ChanSpeechOutput <- buffer.String()

//// And also output the buffer to the file
//_, err = file.Write(buffer.Bytes())
//if err != nil {
//a.ChanError <- err
//close(done)
//}

//close(done)
//} else if a.FileOutput == false {
//// Create a new buffer
//buffer := new(bytes.Buffer)

//// Prepare the speech analysis pipe.
//p := pipe.Line(
//pipe.ReadFile("/home/dzrodrig/dev/projects/src/github.com/ObamaPhony/speech-analysis/test-speeches.txt"),
//pipe.Exec("python3", a.ScriptLocation),
//)

//// Execute the pipe.
//_, err := pipe.CombinedOutput(p)
//if err != nil {
//a.ChanError <- err
//close(done)
//}

//// Pass the buffer to the channel
//speech := buffer.String()
//a.ChanSpeechOutput <- speech

//// Done.
//close(done)

//}
/*}*/

// SpeechAnalysis2 d
func SpeechAnalysis2(a Arguments) (string, error) {
	if a.FileOutput != true {
		// Output to channel only.

		// Create new bytes buffer.
		b := &bytes.Buffer{}
		p := pipe.Line(
			pipe.Read(a.SpeechRequest),
			pipe.Exec(a.ScriptLocation),
			pipe.Tee(b),
		)
		err := pipe.Run(p)
		if err != nil {
			return "error pipe", err
		}

		stringBuffer := &bytes.Buffer{}
		if err := json.Compact(stringBuffer, b.Bytes()); err != nil {
			panic(err)
		}

		return stringBuffer.String(), nil
	}

	// Output to channel and file.

	// Create new bytes buffer.
	b := &bytes.Buffer{}
	p := pipe.Line(
		pipe.Read(a.SpeechRequest),
		pipe.Exec(a.ScriptLocation),
		pipe.Tee(b),
	)
	err := pipe.Run(p)
	if err != nil {
		return "error pipe", err
	}

	stringBuffer := &bytes.Buffer{}
	if err := json.Compact(stringBuffer, b.Bytes()); err != nil {
		panic(err)
	}

	return stringBuffer.String(), nil
}

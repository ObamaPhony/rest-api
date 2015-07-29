package main

import (
	"bytes"
	"fmt"
	"os"
	// "github.com/ObamaPhony/rest-api/controllers"
	obexec "github.com/ObamaPhony/rest-api/exec"
)

func main() {

	go func() {

		cb1 := make(chan *bytes.Buffer)
		ce1 := make(chan error)

		obexec.SpeechAnalysis(cb1, ce1, true)

		x := <-cb1
		y := <-ce1

		if y != nil {
			panic(y)
		}
		result := x.String()
		fmt.Println(result)

		f, err := os.Create("./output.json")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		wr1, err := f.Write(x.Bytes())
		if err != nil {
			panic(err)
		}
		fmt.Printf("Wrote %d bytes\n", wr1)

		f.Sync()

		if err != nil {
			panic(err)
		}

	}()

	// controllers.StartServer(":8080")
}

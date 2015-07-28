package main

import (
	// "fmt"
	"github.com/ObamaPhony/obama-rest-api/controllers"
	// oexec "github.com/ObamaPhony/obama-rest-api/exec"
	// "gopkg.in/pipe.v2"
	// "os"
	// "os/exec"
	// "github.com/ObamaPhony/obama-rest-api/exec"
)

func main() {

	// TODO: Branch off, and refine the async system - possibly create a module.
	// go func() {
	// 	p := pipe.Line(
	// 		pipe.ReadFile("./test-speeches.txt"),
	// 		pipe.Exec("../speech-analysis/analyse-speeches.py"),
	// 	)

	// 	output, err := pipe.CombinedOutput(p)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Printf("%s", output)
	// }()

	controllers.StartServer(":8080")
}

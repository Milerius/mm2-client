package cli

import (
	"fmt"
	"os"
)

func Executor(t string) {
	switch t {
	case "init":
		InitMM2()
	case "exit":
		fmt.Println("Bye")
		os.Exit(0)
	}
	return
}

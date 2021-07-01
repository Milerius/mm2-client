package cli

import (
	"fmt"
	"os"
)

func Executor(t string) {
	switch t {
	case "exit":
		fmt.Println("Bye")
		os.Exit(0)
	}
	return
}

package main

import (
	"os"

	"github.com/sttk/sabi"
)

func main() {
	err := sabi.StartApp(app)

	switch err.Reason().(type) {
	case StdinIsNotTty:
		os.Exit(1)
	case FailToPrint:
		os.Exit(3)
	case InvalidOption:
		os.Exit(2)
	default:
		os.Exit(1)
	}
}

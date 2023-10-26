package main

import (
	"os"

	"github.com/sttk/sabi"
)

func main() {
	if sabi.StartApp(app).IsNotOk() {
		os.Exit(1)
	}
}

package main

import (
	"os"

	"github.com/liampulles/banger/pkg/app"
)

func main() {
	os.Exit(app.Run(os.Args))
}

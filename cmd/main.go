package main

import (
	"io"
	"os"

	"github.com/gokul656/pbcopy-go"
)

func main() {
	clipboard, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	if _, err := pbcopy.Write(clipboard); err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/msrevive/fn-export/cmd"
)

func main() {
	fmt.Println("Starting FN Export tool...")

	if err := cmd.Run(os.Args); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(-1)
	}
}
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting FN Export tool...")

	if err := cmd.Run(os.Args); err != nil {
		panic(err)
	}
}
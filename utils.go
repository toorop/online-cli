package main

import (
	"fmt"
	"os"
)

// dieOnError die on error
func dieOnError(err error) {
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}

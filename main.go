package main

import (
	"log"
	"partial-git/cmd"
)

var Version = "PLACEHOLDER"

func main() {
	// Set the version in the cmd package
	cmd.Version = Version

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

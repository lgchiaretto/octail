package main

import (
	"fmt"
)

var (
	version string
	commit  string
	date    string
)

func printVersion() {
	fmt.Printf("octail version: %s, commit: %s, build at: %s\n", version, commit, date)
}

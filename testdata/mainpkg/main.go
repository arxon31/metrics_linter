package main

import "os"

func main() {
	os.Exit(run()) // want "os.Exit calls are not allowed"
}

func run() int {
	return 1
}

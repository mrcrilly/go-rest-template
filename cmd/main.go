package main

import (
	"github.com/mrcrilly/teehee"
)

func main() {
	teehee.Config("config.toml")
	teehee.StartServer()
}

func checkErrorAndPanic(err error) {
	if err != nil {
		panic(err)
	}
}

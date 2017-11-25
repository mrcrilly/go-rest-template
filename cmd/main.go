package main

import "github.com/mrcrilly/teehee"

func main() {
	teehee.Init("config.toml", "app.log")
	teehee.StartServer()
}

func checkErrorAndPanic(err error) {
	if err != nil {
		panic(err)
	}
}

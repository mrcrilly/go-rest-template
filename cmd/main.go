package main

<<<<<<< 586e99b73b65fc5f47ac168f543b584e03a0f3da
import "github.com/mrcrilly/teehee"

func main() {
	teehee.Init("config.toml")
=======
import (
	"github.com/mrcrilly/teehee"
)

func main() {
	teehee.Config("config.toml")
>>>>>>> wip on config
	teehee.StartServer()
}

func checkErrorAndPanic(err error) {
	if err != nil {
		panic(err)
	}
}

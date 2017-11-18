package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var globalConfig *Configuration
var globalStatus *Status
var globalLogger *log.Logger

func init() {
	globalStatus = new(Status)
	globalStatus.HttpStatusCodes = make(map[int]int, 0)
	globalConfig = new(Configuration)
}

func main() {
	err := globalConfig.Load("config.json")
	checkErrorAndPanic(err)

	//err = buildLogger()
	//checkErrorAndPanic(err)

	err = serveRequests()
	checkErrorAndPanic(err)
}

func checkErrorAndPanic(err error) {
	if err != nil {
		panic(err)
	}
}

//func buildLogger() (err error) {
//fd, err := os.Open(globalConfig.Logging.Path)
//if err != nil {
//if os.IsNotExist(err) {

//}
//}
//globalLogger = log.New(fd, "app: ", log.Lshortfile)
//return
//}

func serveRequests() (err error) {
	router := mux.NewRouter()
	router.HandleFunc("/", HandlerIndex)
	router.HandleFunc("/config", HandlerReadOnlyConfig)
	router.HandleFunc("/health", HandlerHealthCheck)

	server := new(http.Server)

	if globalConfig.Http.Tls {
	} else {
		server.Addr = fmt.Sprintf("%s:%s",
			globalConfig.Http.BindIp, globalConfig.Http.BindPort,
		)
		server.Handler = router
		err = server.ListenAndServe()
	}

	return
}

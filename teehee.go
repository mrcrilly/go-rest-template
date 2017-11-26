package teehee

import (
<<<<<<< 586e99b73b65fc5f47ac168f543b584e03a0f3da
=======
	"io"
>>>>>>> wip on config
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var globalStatus *Status
var globalLogger *logrus.Logger

<<<<<<< 586e99b73b65fc5f47ac168f543b584e03a0f3da
func Init(configFrom string) (err error) {
=======
func config(configFrom io.Reader) (err error) {
>>>>>>> wip on config
	globalStatus = new(Status)
	globalStatus.HttpStatusCodes = make(map[int]int, 0)

	// set some sensible defaults so we can ignore
	// the need for a configuration file
	viper.SetDefault("http.port", "8080")
	viper.SetDefault("http.ip", "127.0.0.1")
	viper.SetDefault("logging.enabled", true)
	viper.SetDefault("logging.file", "app.log")
<<<<<<< 586e99b73b65fc5f47ac168f543b584e03a0f3da

	// attempt to load configuration
	viper.SetConfigName(configFrom)
	viper.AddConfigPath(".")

	// ignoring the error because we've set sensible(?)
	// defaults
	viper.ReadInConfig()

	if viper.GetBool("logging.enabled") {
		globalLogger = logrus.New()
		fd, err := os.OpenFile(viper.GetString("logging.file"),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
=======
	viper.ReadConfig(configFrom)

	if viper.GetBool("logging.enabled") {
		fd, err := os.OpenFile(
			viper.GetString("logging.file"),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0644)
		globalLogger = logrus.New()
>>>>>>> wip on config
		if err != nil {
			return err
		}
		globalLogger.Out = fd
	}

	return
}
<<<<<<< 586e99b73b65fc5f47ac168f543b584e03a0f3da
=======

func ConfigFromReader(configFrom io.Reader) (err error) {
	return config(configFrom)
}

func Config(configFrom string) (err error) {
	fd, err := os.Open(configFrom)
	if err != nil {
		return
	}
	return config(fd)
}
>>>>>>> wip on config

package main

import (
	"encoding/json"
	"errors"
	"os"
)

type Configuration struct {
	Http struct {
		Tls      bool   `json:"tls"`
		BindIp   string `json:"bind_ip"`
		BindPort string `json:"bind_port"`
	} `json:"http"`

	Logging struct {
		File string `json:"file"`
	} `json:"logging"`
}

func (c *Configuration) Load(file string) (err error) {
	if file == "" {
		file = "config.json"
	}

	fi, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("file does not exist")
		}

		return
	}

	if fi.IsDir() {
		return errors.New("cannot load a directory")
	}

	fd, err := os.Open(file)
	if err != nil {
		return
	}

	return json.NewDecoder(fd).Decode(c)
}

package config

import (
	"encoding/json"
	"io"
	"os"
)

const (
	FileName = "config.json"
)

var (
	Global GlobalConfig
)

func Read() {
	file, err := os.Open(FileName)
	if err != nil {
		// TODO: crash
	}
	defer file.Close()

	byteArray, _ := io.ReadAll(file)
	err = json.Unmarshal(byteArray, &Global)
	if err != nil {
		// TODO: crash
	}

	// TODO: log message
}

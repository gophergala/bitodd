package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
)

/*
 * Constants
 */

var (
	configFileLocation string
	fileLock           sync.Mutex
)

type config struct {
	Port       string `json:"port"`
}

var appConfig config

func GetConfig() *config {
	return &appConfig
}

func Load(filename string) error {

	fileLock.Lock()
	defer fileLock.Unlock()

	file, err := os.Open(filename)
	if err != nil {
		return errors.New("Cannot open datafile")
	} else {
		defer file.Close()

		configFileLocation = filename

		decoder := json.NewDecoder(file)
		decErr := decoder.Decode(&appConfig)
		if decErr != nil {
			return errors.New("Error while decoding items: " + decErr.Error())
		}
	}

	log.Println("Loaded config from file")
	return nil
}

func Save() {

	fileLock.Lock()
	defer fileLock.Unlock()

	file, err := os.Create(configFileLocation)
	if err != nil {
		log.Println("Cannot write config file")
		return
	}
	defer file.Close()

	b, err := json.MarshalIndent(appConfig, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}

	_, err = file.Write(b)
	if err != nil {
		log.Println("Error while writing config:", err.Error())
		return
	}
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const defaultConfigFile = "./data/config.json"

// Config config
type Config struct {
	file          string
	TelegramToken string
	NewsAPIToken  string

	StoreFile string

	ChannelID int64
	Country   string

	Debug bool
}

func loadConfig(file string) *Config {
	file = filepath.Clean(file)

	var config Config

	s, err := os.Stat(file)
	if err != nil || s.Size() == 0 {
		createConfig(file)
		log.Println("New config created")
		os.Exit(1)
		return nil
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &config
}

func (config *Config) save() error {
	if config == nil {
		return errors.New("config is nil")
	}

	if len(config.file) == 0 {
		return errors.New("config file is empty")
	}

	data, err := json.MarshalIndent(*config, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(config.file, data, 0600)
}

func createConfig(file string) {
	config := getDefaultConfig(file)
	err := config.save()
	if err != nil {
		log.Fatal(err)
	}
}

func getDefaultConfig(file string) Config {
	return Config{
		file:          file,
		ChannelID:     0,
		NewsAPIToken:  "NewsTokenHere",
		TelegramToken: "TelegramTokenHere",
		Country:       "de",
		StoreFile:     "./data/store.json",
		Debug:         false,
	}
}

func (config *Config) check() bool {
	defalutConfig := getDefaultConfig(config.file)
	success := true

	if config.ChannelID == defalutConfig.ChannelID {
		fmt.Println("You have to set the channel ID!")
		success = false
	}

	if len(config.Country) == 0 {
		fmt.Println("You have to set the country")
		success = false
	}

	if config.NewsAPIToken == defalutConfig.NewsAPIToken {
		fmt.Println("You have to set the news API token")
		success = false
	}

	if config.TelegramToken == defalutConfig.TelegramToken {
		fmt.Println("You have to set the telegram token")
		success = false
	}

	return success
}

func (config *Config) getStorePath() string {
	path := config.StoreFile
	if len(path) == 0 {
		path = getDefaultConfig("").StoreFile
	}

	return filepath.Clean(path)
}

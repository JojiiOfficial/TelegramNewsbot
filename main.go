package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// ConfigEnVar environmentvariable for the config file
const ConfigEnVar = "NB_CONFIG"

func main() {
	config := loadConfig(getConfigFile())
	if config == nil {
		return
	}

	if !config.check() {
		fmt.Println("Config check failed")
		return
	}

	log.Println("Localtime:", time.Now().String())
	log.Println("Refreshinterval:", config.getRefreshInterval())

	bot := newNewsBot(config)
	if bot == nil {
		return
	}

	store := newStore(config.getStorePath())
	if err := store.load(); err != nil {
		fmt.Println(err)
		return
	}

	newsClient := newNewsClient(config)

	for {
		news := newsClient.checkNews(store.LastSync)
		bot.sendNews(news)

		err := store.updateLastSync(time.Now())
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(config.getRefreshInterval())
	}
}

func getConfigFile() string {
	envFile := os.Getenv(ConfigEnVar)
	if len(envFile) > 0 {
		return envFile
	}

	return defaultConfigFile
}

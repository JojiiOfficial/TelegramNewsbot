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

	bot := newNewsBot(config)
	if bot == nil {
		return
	}
	log.Println("Bot started")

	store := newStore(config.getStorePath())
	if err := store.load(); err != nil {
		fmt.Println(err)
		return
	}

	newsClient := newNewsClient(config)
	log.Println("Newsclient connected")

	for {
		news := newsClient.checkNews(store.LastSync)
		bot.sendNews(news)
		time.Sleep(5 * time.Minute)

		err := store.updateLastSync(time.Now())
		if err != nil {
			log.Fatal(err)
		}

	}
}

func getConfigFile() string {
	envFile := os.Getenv(ConfigEnVar)
	if len(envFile) > 0 {
		return envFile
	}

	return defaultConfigFile
}

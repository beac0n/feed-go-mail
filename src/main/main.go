package main

import (
	"encoding/json"
	"feedgomail/src/config"
	"feedgomail/src/feed"
	"log"
	"os"
	"time"
)

func check(err error, reason string) {
	if err != nil {
		log.Fatal("ERROR", reason, err)
	}
}

func main() {
	configFilePath := "./feedgomail.json"
	file, err := os.Open(configFilePath)
	check(err, "ERROR reading config: " + configFilePath)

	decoder := json.NewDecoder(file)
	conf := &config.Config{}
	err = decoder.Decode(&conf)
	check(err, "Error decoding config json:")

	err = file.Close()

	feeds := make([]*feed.Feed, len(conf.Feeds))
	for index, feedUrl := range conf.Feeds {
		feeds[index] = feed.NewFeed(feedUrl, conf)
	}

	for {
		log.Println("sending mails...")
		for _, f := range feeds {
			f.ProcessFeed()
		}

		log.Println("waiting for 10 minutes")
		time.Sleep(10 * time.Minute)
	}
}

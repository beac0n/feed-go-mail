package main

import (
	"encoding/json"
	"feed-go-mail/src/config"
	"feed-go-mail/src/feed"
	"flag"
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
	configFilePath := flag.String("config-path", "./feedgomail.json", "path to config file")

	flag.Parse()

	file, err := os.Open(*configFilePath)
	check(err, "ERROR reading config: " + *configFilePath)

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

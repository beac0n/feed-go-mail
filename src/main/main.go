package main

import (
	"flag"
	"gofeedtomail/src/config"
	"gofeedtomail/src/feed"
	"log"
	"time"
)

type feedList []string

func (feedList *feedList) String() string {
	return ""
}

func (feedList *feedList) Set(value string) error {
	*feedList = append(*feedList, value)
	return nil
}

func main() {

	from := flag.String("from", "", "email sender")
	to := flag.String("to", "", "email receiver")
	host := flag.String("host", "", "smtp host")
	port := flag.Int64("port", 0, "smtp port")
	password := flag.String("password", "", "smtp password")

	var feedList feedList
	flag.Var(&feedList, "feeds", "list of feeds to crawl")

	flag.Parse()

	conf := &config.Config{
		From:     *from,
		To:       *to,
		Host:     *host,
		Port:     *port,
		Password: *password,
		Feeds:    feedList,
	}

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

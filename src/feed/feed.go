package feed

import (
	"github.com/mmcdole/gofeed"
	"gofeedtomail/src/config"
	"log"
	"net/smtp"
	"strconv"
	"time"
)

type Feed struct {
	url                string
	latestSentItemLink string
	parser             *gofeed.Parser
	conf               *config.Config
}

func (feed *Feed) getIndexOfLatestSentItem(items []*gofeed.Item) int {
	defer feed.updateLatestSentItemLink(items)

	if feed.latestSentItemLink == "" {
		return 0
	}

	for index, item := range items {
		if feed.latestSentItemLink == item.Link {
			return index
		}
	}

	return 0
}

func (feed *Feed) updateLatestSentItemLink(items []*gofeed.Item) {
	if len(items) > 0 {
		feed.latestSentItemLink = items[0].Link
	}
}

var CRLF = "\r\n"
var newLine = "<br/><br/>"

func (feed *Feed) processFeedItem(feedTitle string, item *gofeed.Item) {
	msg := "From: " + feed.conf.From + CRLF +
		"To: " + feed.conf.To + CRLF +
		"Subject: " + feedTitle + ": " + item.Title + CRLF +
		"MIME-Version: 1.0" + CRLF +
		"Content-Type: text/html;charset=UTF-8" + CRLF +
		CRLF + // email body starts below here
		`<a href="` + item.Link + `">` + item.Link + "</a>" + newLine +
		item.Description + newLine +
		item.Content

	log.Println("sending mail for " + item.Link)
	address := feed.conf.Host + ":" + strconv.FormatInt(feed.conf.Port, 10)
	auth := smtp.PlainAuth("", feed.conf.From, feed.conf.Password, feed.conf.Host)
	if err := smtp.SendMail(address, auth, feed.conf.From, []string{feed.conf.To}, []byte(msg)); err != nil {
		log.Fatal(err)
	}
}

func (feed *Feed) ProcessFeed() {
	parsedFeed, err := feed.parser.ParseURL(feed.url)
	if err != nil {
		log.Fatal(err)
		return
	}

	if len(parsedFeed.Items) == 0 {
		return
	}

	log.Println("sending mails for url " + parsedFeed.Link)
	for i := feed.getIndexOfLatestSentItem(parsedFeed.Items); i > 0; i-- {
		feed.processFeedItem(parsedFeed.Title, parsedFeed.Items[i-1])
		// wait a second before sending the next email
		time.Sleep(time.Second)
	}
}

func NewFeed(feedUrl string, conf *config.Config) *Feed {
	return &Feed{url: feedUrl, parser: gofeed.NewParser(), conf: conf}
}

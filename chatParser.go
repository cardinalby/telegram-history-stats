package main

import (
	"github.com/PuerkitoBio/goquery"
	"os"
	"strings"
	"time"
)

type chatMessage struct {
	fromMe   bool
	fromName string
	text     string
	time     time.Time
}

func getChatName(files *chatFiles) (string, error) {
	file, _ := os.Open(files.files[0])
	defer func() { _ = file.Close() }()

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		return "", err
	}
	selection := doc.Find(".page_wrap .page_header a div.bold.text")
	return strings.TrimSpace(selection.First().Text()), nil
}

func readChatFilesMessages(files *chatFiles, done <-chan struct{}) chan *chatMessage {
	allMessages := make(chan *chatMessage)

	go func() {
		defer close(allMessages)

		for _, file := range files.files {
			for message := range readFileMessages(file, done) {
				select {
				case allMessages <- message:
				case <-done:
					return
				}

			}
		}
	}()

	return allMessages
}

func readFileMessages(filePath string, done <-chan struct{}) chan *chatMessage {
	file, _ := os.Open(filePath)
	doc, err := goquery.NewDocumentFromReader(file)
	messages := make(chan *chatMessage)
	if err != nil {
		_ = file.Close()
		return nil
	}
	chatName := strings.TrimSpace(doc.Find(".page_wrap .page_header a div.bold.text").First().Text())

	go func() {
		defer func() { _ = file.Close() }()
		defer close(messages)

		doc.Find("div.message.default.clearfix").EachWithBreak(func(i int, msgTag *goquery.Selection) bool {
			timeStr, exists := msgTag.Find("div.date.details").First().Attr("title")
			if !exists {
				return true
			}
			msgTime, err := time.Parse("02.01.2006 15:04:05", timeStr)
			if err != nil {
				return true
			}
			fromName := strings.TrimSpace(msgTag.Find("div.from_name").First().Text())
			messageText := msgTag.Find("div.text").First().Text()
			message := &chatMessage{
				fromMe:   fromName != chatName,
				fromName: fromName,
				text:     messageText,
				time:     msgTime,
			}

			select {
			case messages <- message:
				return true
			case <-done:
				return false
			}
		})
	}()

	return messages
}

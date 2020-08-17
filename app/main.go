package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Arg 1: path to folder with exported history")
		fmt.Println("Arg 2 (optional): interval in hours to split sessions. Default: 10")
		return
	}
	historyDir := os.Args[1]

	splitDuration := time.Hour * 10
	if len(os.Args) > 2 {
		hours, err := strconv.Atoi(os.Args[2])
		if err != nil {
			panic("Second arg (split interval, hours) is not integer")
		}
		splitDuration = time.Hour * time.Duration(hours)
	}

	outCsvPath := path.Join(historyDir, "stats.csv")

	done := make(chan struct{})
	chatFilesChan, errors := getChatFiles(historyDir, done)

	csvExporter, err := createCsvExporter(outCsvPath)
	if err != nil {
		panic(err)
	}
	defer csvExporter.Close()

	for {
		select {
		case chat := <-chatFilesChan:
			if chat == nil {
				return
			}
			chatName, err := getChatName(chat)
			if err != nil {
				return
			}

			sessions := splitIntoSessions(splitDuration, readChatFilesMessages(chat, done), done)
			stats := calculateStats(sessions)
			if stats.messagesFromMe+stats.messagesFromContact > 0 {
				fmt.Printf("%s: %d messages", chatName, stats.messagesFromMe+stats.messagesFromContact)
				fmt.Println("")
				_ = csvExporter.WriteRecord(chatName, stats)
			}

		case err := <-errors:
			if err == nil {
				return
			}
			fmt.Println(err)
		case <-done:
			return
		}
	}
}

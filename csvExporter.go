package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type csvExporter struct {
	file   *os.File
	writer *csv.Writer
}

func createCsvExporter(filePath string) (*csvExporter, error) {
	file, err := os.Create(filePath)
	exporter := &csvExporter{
		file:   file,
		writer: nil,
	}
	if err != nil {
		return exporter, nil
	}
	exporter.writer = csv.NewWriter(file)
	err = exporter.writer.Write([]string{
		"Contact",
		"I initiated %",
		"My msgs/session",
		"Contact's msgs/session",
		"My chars/message",
		"Contact's chars/message",
		"Avg messages/session",
		"My avg reaction (sec)",
		"Contact's avg reaction (sec)",
	})
	if err != nil {
		return exporter, err
	}
	return exporter, nil
}

func (exporter csvExporter) WriteRecord(chatName string, stats *chatStats) error {
	totalSessions := stats.sessionsStartedByMe + stats.sessionsStartedByContact
	totalMessages := stats.messagesFromMe + stats.messagesFromContact

	var myReactionTime string
	if stats.myReactionTime.messages > 0 {
		myReactionTime = fmt.Sprintf("%.0f", stats.myReactionTime.Avg().Seconds())
	} else {
		myReactionTime = ""
	}

	var contactReactionTime string
	if stats.contactReactionTime.messages > 0 {
		contactReactionTime = fmt.Sprintf("%.0f", stats.contactReactionTime.Avg().Seconds())
	} else {
		contactReactionTime = ""
	}

	row := []string{
		chatName,
		fmt.Sprintf("%0.f", float32(stats.sessionsStartedByMe)/float32(totalSessions)*100),
		fmt.Sprintf("%.2f", float32(stats.messagesFromMe)/float32(totalSessions)),
		fmt.Sprintf("%.2f", float32(stats.messagesFromContact)/float32(totalSessions)),
		fmt.Sprintf("%.2f", float32(stats.charsFromMe)/float32(stats.messagesFromMe)),
		fmt.Sprintf("%.2f", float32(stats.charsFromContact)/float32(stats.messagesFromContact)),
		fmt.Sprintf("%.2f", float32(totalMessages)/float32(totalSessions)),
		myReactionTime,
		contactReactionTime,
	}

	return exporter.writer.Write(row)
}

func (exporter csvExporter) Close() {
	if exporter.writer != nil {
		exporter.writer.Flush()
	}
	if exporter.file != nil {
		_ = exporter.file.Close()
	}
}

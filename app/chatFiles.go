package main

import (
	"io/ioutil"
	"path"
)

type chatFiles struct {
	chatDir string
	files   []string
}

func getChatFiles(historyDir string, done <-chan struct{}) (<-chan *chatFiles, <-chan error) {
	result := make(chan *chatFiles)
	errors := make(chan error)

	chatsDirPath := path.Join(historyDir, "chats")
	chats, err := ioutil.ReadDir(chatsDirPath)
	go func() {
		defer close(result)
		defer close(errors)

		if err != nil {
			errors <- err
			return
		}

		for _, dir := range chats {
			if !dir.IsDir() {
				continue
			}
			chatDirPath := path.Join(chatsDirPath, dir.Name())
			chatFilesInfo, err := ioutil.ReadDir(chatDirPath)
			if err != nil {
				select {
				case errors <- err:
					break
				case <-done:
					return
				}
			}
			resultItem := chatFiles{
				chatDir: chatDirPath,
				files:   []string{},
			}
			for _, file := range chatFilesInfo {
				if !file.IsDir() {
					resultItem.files = append(resultItem.files, path.Join(chatDirPath, file.Name()))
				}
			}
			select {
			case result <- &resultItem:
				continue
			case <-done:
				return
			}
		}
	}()

	return result, errors
}

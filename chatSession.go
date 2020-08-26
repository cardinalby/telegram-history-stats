package main

import "time"

type chatSession struct {
	myMessages          int
	contactMessages     int
	myChars             int
	contactChars        int
	startedAt           time.Time
	finishedAt          time.Time
	initiatedByMe       bool
	myReactionTime      messageReactionTime
	contactReactionTime messageReactionTime
}

func splitIntoSessions(
	splitDuration time.Duration,
	messages chan *chatMessage,
	done chan struct{},
) chan *chatSession {
	sessions := make(chan *chatSession)

	go func() {
		defer close(sessions)

		var session *chatSession
		var prevMessage *chatMessage
		for message := range messages {
			// If should create new session
			if session == nil || message.time.Sub(session.finishedAt) > splitDuration {
				if session != nil {
					select {
					case sessions <- session:
					case <-done:
						return
					}
				}
				session = &chatSession{
					startedAt:     message.time,
					initiatedByMe: message.fromMe,
				}
			} else if prevMessage != nil && message.fromMe != prevMessage.fromMe {
				if message.fromMe {
					session.myReactionTime.AddMessageReaction(message.time.Sub(prevMessage.time))
				} else {
					session.contactReactionTime.AddMessageReaction(message.time.Sub(prevMessage.time))
				}
			}

			session.finishedAt = message.time
			if message.fromMe {
				session.myChars += len(message.text)
				session.myMessages++
			} else {
				session.contactChars += len(message.text)
				session.contactMessages++
			}
			prevMessage = message
		}

		if session != nil {
			sessions <- session
		}
	}()

	return sessions
}

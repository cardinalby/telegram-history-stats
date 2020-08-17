package main

import "time"

type messageReactionTime struct {
	messages      int
	totalDuration time.Duration
}

func (rt *messageReactionTime) Avg() time.Duration {
	return rt.totalDuration / time.Duration(rt.messages)
}

func (rt *messageReactionTime) Add(reactionTime messageReactionTime) messageReactionTime {
	return messageReactionTime{
		messages:      rt.messages + reactionTime.messages,
		totalDuration: rt.totalDuration + reactionTime.totalDuration,
	}
}

func (rt *messageReactionTime) AddMessageReaction(reactionTime time.Duration) {
	rt.messages++
	rt.totalDuration += reactionTime
}

package main

type chatStats struct {
	sessionsStartedByMe      int64
	sessionsStartedByContact int64
	messagesFromMe           int64
	messagesFromContact      int64
	charsFromMe              int64
	charsFromContact         int64
	avgMessagesPerSession    float32
	myReactionTime           messageReactionTime
	contactReactionTime      messageReactionTime
}

func calculateStats(sessions chan *chatSession) *chatStats {
	var stats chatStats

	for session := range sessions {
		stats.messagesFromMe += int64(session.myMessages)
		stats.messagesFromContact += int64(session.contactMessages)
		stats.charsFromMe += int64(session.myChars)
		stats.charsFromContact += int64(session.contactChars)
		stats.myReactionTime = stats.myReactionTime.Add(session.myReactionTime)
		stats.contactReactionTime = stats.contactReactionTime.Add(session.contactReactionTime)

		if session.initiatedByMe {
			stats.sessionsStartedByMe++
		} else {
			stats.sessionsStartedByContact++
		}
	}

	stats.avgMessagesPerSession = float32(stats.messagesFromMe+stats.messagesFromContact) /
		float32(stats.sessionsStartedByMe+stats.sessionsStartedByContact)

	return &stats
}

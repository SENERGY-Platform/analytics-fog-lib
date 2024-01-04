package upstream

type UpstreamControlMessage struct {
	OperatorOutputTopic string `json:"topic"`
}

type UpstreamSyncMessage struct {
	OperatorOutputTopics []string `json:"topics"`
}
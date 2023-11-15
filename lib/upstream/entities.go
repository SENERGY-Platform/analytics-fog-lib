package upstream

type UpstreamEnableMessage struct {
	OperatorOutputTopic string `json:"topic"`
}

type UpstreamDisableMessage UpstreamEnableMessage

type UpstreamEnvelope struct {
	BaseOperatorName string `json:"baseOperatorName"`
	OperatorID string `json:"operatorId"`
	Message string `json:"message"`
}
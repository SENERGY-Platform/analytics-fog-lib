package operator

import (
	"github.com/SENERGY-Platform/analytics-fog-lib/lib/topic"
)

const OperatorsTopicPrefix = topic.TopicPrefix + "operators/"

// Used by the connector to send control messages like start/stop operator to the master.
const OperatorsControlTopic = OperatorsTopicPrefix + "control"

// Used by fog upstream proxy to publish all messages that must be forwarded to the cloud by the future connector 
const OperatorsResultTopic = OperatorsTopicPrefix + "results"

// Used by the agents to inform master about success/erros of operators
const OperatorsControlResponseTopic = OperatorsTopicPrefix + "control/response"
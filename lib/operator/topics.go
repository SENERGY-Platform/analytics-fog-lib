package operator

import (
	"github.com/SENERGY-Platform/analytics-fog-lib/lib/topic"
)

const OperatorsTopicPrefix = topic.TopicPrefix + "operators/"

// Used by the connector to send control messages like start/stop operator to the master.
const OperatorsControlTopic = OperatorsTopicPrefix + "control"

// Used by the last operators in a pipeline to publish results that are forwarded to platform
const OperatorsResultTopic = OperatorsTopicPrefix + "results"

// Used by the agents to inform master about success/erros of operators
const OperatorsControlResponseTopic = OperatorsTopicPrefix + "control/response"
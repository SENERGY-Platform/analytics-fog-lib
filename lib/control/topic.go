package control

import (
	"github.com/SENERGY-Platform/analytics-fog-lib/lib/topic"
)

// Used by the connector to send control messages like start/stop operator to the master.
// Used by the Master to send control messages to agents
const ControlTopic = topic.TopicPrefix + "control"

func GetConnectorControlTopic(userID string) string {
	// Used by the Connector to receive user specific control messages from the platform
	return topic.TopicPrefix + "control/" + userID
}

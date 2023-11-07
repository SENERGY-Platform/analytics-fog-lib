package control

import (
	"github.com/SENERGY-Platform/analytics-fog-lib/lib/topic"
)

func GetConnectorControlTopic(userID string) string {
	// Used by the Connector to receive user specific control messages from the platform
	return topic.TopicPrefix + "control/" + userID
}

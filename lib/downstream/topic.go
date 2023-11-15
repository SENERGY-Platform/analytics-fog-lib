package downstream

import (
	"github.com/SENERGY-Platform/analytics-fog-lib/lib/control"
)

// Used by kafka2mqtt to send cloud operator results or device data downstream to fog environment  
const DownstreamProxyTopic = "/downstream" 

func GetDownstreamTopic(userID string) string {
	return control.GetConnectorControlTopic(userID) + DownstreamProxyTopic
}


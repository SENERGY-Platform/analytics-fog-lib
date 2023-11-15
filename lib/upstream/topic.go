package upstream

import (
	"github.com/SENERGY-Platform/analytics-fog-lib/lib/control"
)

// Used by the connector/upstream proxy to forward operator output to platform
const UpstreamProxyTopic = "/upstream" 

const UpstreamProxyEnableTopic = UpstreamProxyTopic + "/enable"
const UpstreamProxyDisableTopic = UpstreamProxyTopic + "/disable"

func GetUpstreamProxyEnableTopic(userID string) string {
	return control.GetConnectorControlTopic(userID) + UpstreamProxyTopic + UpstreamProxyEnableTopic
}

func GetUpstreamProxyDisableTopic(userID string) string {
	return control.GetConnectorControlTopic(userID) + UpstreamProxyTopic + UpstreamProxyDisableTopic
}

package downstream

import (
	"github.com/SENERGY-Platform/analytics-fog-lib/lib/cloud"
)

// Used by kafka2mqtt to send cloud operator results or device data downstream to fog environment  
const DownstreamTopic = "/downstream"

func GetDownstreamOperatorCloudTopic(id string) string {
	return cloud.CloudAnalyticsMQTTTopicPrefix + id + DownstreamTopic + "/operator/#"
}

func GetDownstreamDeviceCloudTopic(id string) string {
	return cloud.CloudAnalyticsMQTTTopicPrefix + id + DownstreamTopic + "/device/#"
}

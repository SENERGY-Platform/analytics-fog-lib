package downstream

import (
	"github.com/SENERGY-Platform/analytics-fog-lib/lib/cloud"
)

// Used by kafka2mqtt to send cloud operator results or device data downstream to fog environment  
const DownstreamTopic = "/downstream"

func GetDownstreamOperatorCloudSubTopic(id string) string {
	return GetDownstreamOperatorCloudMatchTopic(id) + "/#"
}

func GetDownstreamOperatorCloudMatchTopic(id string) string {
	return cloud.CloudAnalyticsMQTTTopicPrefix + id + DownstreamTopic + "/operator"
}

func GetDownstreamDeviceCloudSubTopic(id string) string {
	return cloud.CloudAnalyticsMQTTTopicPrefix + id + DownstreamTopic + "/device/#"
}

func GetDownstreamOperatorCloudPubTopicPrefix(id string) string {
	return cloud.CloudAnalyticsMQTTTopicPrefix + id + DownstreamTopic + "/operator/"
}

func GetDownstreamDeviceCloudPubTopicPrefix(id string) string {
	return cloud.CloudAnalyticsMQTTTopicPrefix + id + DownstreamTopic + "/device/"
}

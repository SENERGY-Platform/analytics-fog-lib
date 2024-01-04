package upstream

import (
	"github.com/SENERGY-Platform/analytics-fog-lib/lib/cloud"	
	"strings"

)

// Used by the connector/upstream proxy to forward operator output to platform
const UpstreamTopic = "upstream" 

func GetUpstreamEnableCloudTopic(id string) string {
	return cloud.CloudAnalyticsMQTTTopicPrefix + id + "/" + UpstreamTopic + "/enable"
} 

func GetUpstreamDisableCloudTopic(id string) string {
	return cloud.CloudAnalyticsMQTTTopicPrefix + id + "/" + UpstreamTopic + "/disable"
} 

// Used by fog connector to publish upstream messages
// Operator Name is appended to enable easy mapping to kafka at cloud side
const CloudUpstreamTopic = cloud.CloudAnalyticsMQTTTopicPrefix + UpstreamTopic + "/messages"

func ParseCloudUpstreamTopic(topic string) string {
	splittedTopic := strings.Split(topic, "/")
	return splittedTopic[len(splittedTopic)-1]
}

func GetUpstreamControlSyncResponseTopic(id string) string {
	return cloud.CloudAnalyticsMQTTTopicPrefix + id + "/" + UpstreamTopic + "/sync/response"
}

func GetUpstreamControlSyncTriggerPubTopic(id string) string {
	return cloud.CloudAnalyticsMQTTTopicPrefix + id + "/" + UpstreamTopic + "/sync/request"
}

func GetUpstreamControlSyncTriggerSubTopic() string {
	return cloud.CloudAnalyticsMQTTTopicPrefix + "#/" + UpstreamTopic + "/sync/request"
}

func GetUserIDFromUpstreamControlSyncTopic(topic string) string {
	parts := strings.Split(topic, "/")
	return parts[2]
} 
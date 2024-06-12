package operator

import (
	"github.com/SENERGY-Platform/analytics-fog-lib/lib/fog"
	"fmt"
	"strings"
	
	"github.com/SENERGY-Platform/analytics-fog-lib/lib/cloud"

)

// FOG
const OperatorsTopicPrefix = fog.FogAnalyticsTopicPrefix + "operator/"

// Used by the connector to send control messages like start/stop operator to the master.
const StartOperatorFogTopic = OperatorsTopicPrefix + "control/start"
const StopOperatorFogTopic = OperatorsTopicPrefix + "control/stop"

// Used by the agents to inform master about success/erros of operators
const StartOperatorResponseFogTopic = OperatorsTopicPrefix + "control/start/response"
const StopOperatorResponseFogTopic = OperatorsTopicPrefix + "control/stop/response"

// Used by fog master to get sync 
const OperatorControlSyncResponseFogTopic = OperatorsTopicPrefix + "control/sync/response"

// CLOUD 
func GetStartOperatorCloudTopic(id string) string {
	return cloud.CloudAnalyticsMQTTTopicPrefix + id + "/operator/control/start"
}

func GetStopOperatorCloudTopic(id string) string {
	return cloud.CloudAnalyticsMQTTTopicPrefix + id + "/operator/control/stop"
}

func GetOperatorControlSyncResponseTopic(id string) string {
	return cloud.CloudAnalyticsMQTTTopicPrefix + id + "/operator/control/sync/response"
}

func GetOperatorControlSyncTriggerPubTopic(id string) string {
	return cloud.CloudAnalyticsMQTTTopicPrefix + id + "/operator/control/sync/request"
}

func GetOperatorControlSyncTriggerSubTopic() string {
	return cloud.CloudAnalyticsMQTTTopicPrefix + "+/operator/control/sync/request"
}

func GetUserIDFromOperatorControlSyncTopic(topic string) string {
	parts := strings.Split(topic, "/")
	return parts[2]
} 
	
func GenerateFogOperatorTopic(name, operatorID, pipelineID string) string {
	return fmt.Sprintf("%s%s/%s/%s", fog.FogAnalyticsTopicPrefix, name, operatorID, pipelineID)
}

func GenerateCloudOperatorTopic(name string) string {
	return fmt.Sprintf("%s%s", cloud.CloudAnalyticsKafkaTopicPrefix, name)
}
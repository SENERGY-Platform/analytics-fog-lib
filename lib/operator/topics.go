package operator

import (
	"github.com/SENERGY-Platform/analytics-fog-lib/lib/fog"
	"errors"
	"fmt"
	"strings"
	
	"github.com/SENERGY-Platform/analytics-fog-lib/lib/location"
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
	

func GenerateOperatorOutputTopic(name string, baseOperatorID string, operatorID string, deployLocation string) (string, error) {
	if deployLocation == location.Cloud {
		return cloud.CloudAnalyticsKafkaTopicPrefix + name, nil
	} else if deployLocation == location.Local {
		return fog.FogAnalyticsTopicPrefix + name + "/" + operatorID, nil
	}
	return "", errors.New(fmt.Sprintf("Deployment location %s is not known", deployLocation))
}
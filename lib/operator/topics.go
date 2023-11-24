package operator

import (
	"github.com/SENERGY-Platform/analytics-fog-lib/lib/fog"
	"errors"
	"fmt"
	
	"github.com/SENERGY-Platform/analytics-fog-lib/lib/location"
	"github.com/SENERGY-Platform/analytics-fog-lib/lib/cloud"

)

const OperatorsTopicPrefix = fog.FogAnalyticsTopicPrefix + "operator/"

// Used by the connector to send control messages like start/stop operator to the master.
const OperatorsControlFogTopic = OperatorsTopicPrefix + "control"

// Used by the agents to inform master about success/erros of operators
const OperatorsControlResponseFogTopic = OperatorsTopicPrefix + "control/response"

func GetOperatorControlCloudTopic(id string) string {
	return cloud.CloudAnalyticsMQTTTopicPrefix + id + "/operator/control"
}

func GenerateOperatorOutputTopic(name string, baseOperatorID string, operatorID string, deployLocation string) (string, error) {
	if deployLocation == location.Cloud {
		return cloud.CloudAnalyticsKafkaTopicPrefix + name, nil
	} else if deployLocation == location.Local {
		return fog.FogAnalyticsTopicPrefix + name + "/" + operatorID, nil
	}
	return "", errors.New(fmt.Sprintf("Deployment location %s is not known", deployLocation))
}


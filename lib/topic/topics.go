package topic

import (
	"errors"
	"fmt"

	"github.com/SENERGY-Platform/analytics-fog-lib/lib/location"
)

// Used at the fog and platform broker 
const TopicPrefix = "fog/analytics/"

const CloudTopicPrefix = "analytics-"

func GenerateOperatorOutputTopic(name string, baseOperatorID string, operatorID string, deployLocation string) (string, error) {
	if deployLocation == location.Cloud {
		return CloudTopicPrefix + name, nil
	} else if deployLocation == location.Local {
		return TopicPrefix + name + "/" + operatorID, nil
	}
	return "", errors.New(fmt.Sprintf("Deployment location %s is not known", deployLocation))
}

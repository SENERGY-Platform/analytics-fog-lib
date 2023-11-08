package topic

import (
	"errors"
	"fmt"

	"github.com/SENERGY-Platform/analytics-fog-lib/lib/location"
)

const TopicPrefix = "fog/analytics/"

const CloudTopicPrefix = "analytics-"

func GenerateOperatorOutputTopic(name string, deployLocation string) (string, error) {
	if deployLocation == location.Cloud {
		return CloudTopicPrefix + name, nil
	} else if deployLocation == location.Local {
		return TopicPrefix + name, nil
	}
	return "", errors.New(fmt.Sprintf("Deployment location %s is not known", deployLocation))
}

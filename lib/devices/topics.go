package devices

import (
	"errors"

	"github.com/SENERGY-Platform/analytics-fog-lib/lib/location"
)

const LOCAL_EVENT_TOPIC = "event/"
func GetDeviceOutputTopic(deviceID, serviceID, deploymentLocation string) (string, error) {
	if(deploymentLocation == location.Local) {
		topic := LOCAL_EVENT_TOPIC + deviceID + "/" + serviceID
		return topic, nil
	} else if deploymentLocation == location.Cloud {
		return serviceID, nil
	}
	return "", errors.New("Deployment Location not supported: " + deploymentLocation)
}
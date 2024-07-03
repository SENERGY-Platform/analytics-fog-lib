package devices

const LOCAL_EVENT_TOPIC = "event/"
func GetLocalDeviceOutputTopic(deviceLocalID, serviceName string) (string) {
	return LOCAL_EVENT_TOPIC + deviceLocalID + "/" + serviceName
}
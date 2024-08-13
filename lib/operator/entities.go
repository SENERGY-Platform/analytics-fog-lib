package operator

import "time"

type StartOperatorControlCommand struct {
	ImageId        string            `json:"imageId"`
	OperatorConfig map[string]string `json:"operatorConfig"`
	InputTopics    []InputTopic      `json:"inputTopics"`
	OutputTopic string `json:"outputTopic"`
	OperatorIDs
}

type StartOperatorAgentResponse struct {
	OperatorAgentResponse
	ContainerId string `json:"containerId"`
}

type StopOperatorControlCommand struct {
	OperatorIDs
}

type StopOperatorAgentControlCommand struct {
	ContainerId string `json:"deployment_ref"`
	OperatorIDs
}

type StopOperatorAgentResponse struct {
	OperatorAgentResponse
}

type OperatorAgentResponse struct {
	OperatorIDs
	Error string              `json:"error"`
	DeploymentState string `json:"state"`
	AgentId           string `json:"agent_id"`
	Time time.Time `json:"time"`
}

type Operator struct {
	StartOperatorControlCommand
	DeploymentState string                `json:"state"`
	DeploymentError string `json:"deployment_error"`
	AgentId string                `json:"agent_id"`
	ContainerId string `json:"container_id"`
	TimeOfLastHeartbeat time.Time `json:"time_last_heartbeat"`
}

type OperatorIDs struct {
	PipelineId     string `json:"pipelineId"`
	OperatorId     string `json:"operatorId"` // TODO check if operator is unique
	BaseOperatorId string `json:"baseOperatorId"`
}

type InputTopic struct {
	Name        string    `json:"name"`
	FilterType  string    `json:"filterType"`
	FilterValue string    `json:"filterValue"`
	Mappings    []Mapping `json:"mappings"`
}

type Mapping struct {
	Dest   string `json:"dest"`
	Source string `json:"source"`
}
package operator

import (
	"github.com/SENERGY-Platform/analytics-fog-lib/lib/agent"
	"github.com/SENERGY-Platform/analytics-fog-lib/lib/control"
)

type StartOperatorControlCommand struct {
	control.ControlMessage
	Operator StartOperatorMessage `json:"data,omitempty"`
}

type StopOperatorControlCommand struct {
	control.ControlMessage
	OperatorIDs
}

type StartOperatorMessage struct {
	ImageId        string            `json:"imageId"`
	OperatorConfig map[string]string `json:"operatorConfig"`
	InputTopics    []InputTopic      `json:"inputTopics"`
	Config         FogConfig         `json:"config"`
}

type OperatorAgentResponse struct {
	Response        string              `json:"state"`
	ResponseMessage string              `json:"responseMessage"`
	OperatorId      string              `json:"operatorId"`
	Agent           agent.Configuration `json:"agent"`
}

type OperatorAgentSuccessResponse struct {
	OperatorAgentResponse
	ContainerId string `json:"containerId"`
}

type Operator struct {
	StartOperatorMessage
	Event OperatorAgentResponse `json:"event"`
	State string                `json:"state"`
	Agent string                `json:"agent_id"`
}

type OperatorIDs struct {
	PipelineId     string `json:"pipelineId"`
	OperatorId     string `json:"operatorId"` // TODO check if operator is unique
	BaseOperatorId string `json:"baseOperatorId"`
}

type FogConfig struct {
	OutputTopic string `json:"outputTopic"`
	OperatorIDs
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

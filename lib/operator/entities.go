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
	OperatorId string `json:"operatorId"`
}

type StartOperatorMessage struct {
	ImageId        string            `json:"imageId"`
	OperatorConfig map[string]string `json:"operatorConfig"`
	InputTopics    []InputTopic      `json:"inputTopics"`
	Config         FogConfig         `json:"config"`
}

type StartOperatorAgentResponse struct {
	Response        string `json:"response"`
	ResponseMessage string `json:"responseMessage"`
}

type StartOperatorAgentSuccessResponse struct {
	StartOperatorAgentResponse
	OperatorId  string `json:"operatorId"`
	ContainerId string `json:"containerId"`
}

type OperatorJob struct {
	StartOperatorMessage
	StartOperatorAgentResponse

	// Set by the master for the agent
	Agent agent.Agent `json:"agent,omitempty"`
}

type FogConfig struct {
	PipelineId  string `json:"pipelineId,omitempty"`
	OutputTopic string `json:"outputTopic,omitempty"`
	// TODO check if operator is unique
	OperatorId     string `json:"operatorId,omitempty"`
	BaseOperatorId string `json:"baseOperatorId,omitempty"`
}

type InputTopic struct {
	Name        string    `json:"name,omitempty"`
	FilterType  string    `json:"filterType,omitempty"`
	FilterValue string    `json:"filterValue,omitempty"`
	Mappings    []Mapping `json:"mappings,omitempty"`
}

type Mapping struct {
	Dest   string `json:"dest,omitempty"`
	Source string `json:"source,omitempty"`
}

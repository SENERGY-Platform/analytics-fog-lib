package operator

import (
	"github.com/SENERGY-Platform/analytics-fog-lib/lib/agent"
)

type StartOperatorControlCommand struct {
	ImageId        string            `json:"imageId"`
	OperatorConfig map[string]string `json:"operatorConfig"`
	InputTopics    []InputTopic      `json:"inputTopics"`
	Config         FogConfig         `json:"config"`
}

type StartOperatorAgentResponse struct {
	OperatorAgentResponse
	ContainerId string `json:"containerId"`
}

type StopOperatorControlCommand struct {
	OperatorIDs
}

type StopOperatorAgentControlCommand struct {
	DeploymentReference string `json:"deployment_ref"`
	OperatorID string `json:"operator_id"`
}

type StopOperatorAgentResponse struct {
	OperatorAgentResponse
}

type OperatorAgentResponse struct {
	Success        bool              `json:"success"`
	Error string              `json:"error"`
	OperatorState string `json:"state"`
	OperatorId      string              `json:"operatorId"`
	Agent           agent.Configuration `json:"agent"`
}

type Operator struct {
	StartOperatorControlCommand
	Event OperatorAgentResponse `json:"event"`
	State string                `json:"state"`
	Agent string                `json:"agent_id"`
	DeploymentReference string `json:"deployment_ref"`
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
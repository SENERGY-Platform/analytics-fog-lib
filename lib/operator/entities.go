package operator

import (
	"github.com/SENERGY-Platform/analytics-fog-lib/lib/control"
)

type OperatorControlCommand struct {
	control.ControlCommand
	Data OperatorJob `json:"data,omitempty"`
}

type OperatorJob struct {
	ImageId        string            `json:"imageId,omitempty"`
	OperatorConfig map[string]string `json:"operatorConfig,omitempty"`
	InputTopics    []InputTopic      `json:"inputTopics,omitempty"`
	Config         FogConfig         `json:"config,omitempty"`

	// Set by the master for the agent
	Agent Agent `json:"agent,omitempty"`

	// Come from the agent reseponse
	ContainerId     string `json:"containerId,omitempty"`
	Response        string `json:"response,omitempty"`
	ResponseMessage string `json:"responseMessage,omitempty"`
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

package pipeline 

import (
	"github.com/google/uuid"
)

type Pipeline struct {
	Id                 uuid.UUID  `json:"id,omitempty"`
	FlowId             string     `json:"flowId,omitempty"`
	Name               string     `json:"name,omitempty"`
	Description        string     `json:"description,omitempty"`
	Image              string     `json:"image,omitempty"`
	WindowTime         int        `json:"windowTime,omitempty"`
	MergeStrategy      string     `json:"mergeStrategy,omitempty"`
	ConsumeAllMessages bool       `json:"consumeAllMessages,omitempty"`
	Metrics            bool       `json:"metrics,omitempty"`
	MetricsData        Metrics    `json:"metricsData,omitempty"`
	Operators          []Operator `json:"operators,omitempty"`
}

type UpstreamConfig struct {
	Enabled bool 
}

type DownstreamConfig struct {
	Enabled bool
	InstanceID string 
}

type Operator struct {
	Id              string            `json:"id,omitempty"`
	Name            string            `json:"name,omitempty"`
	ApplicationId   uuid.UUID         `json:"applicationId,omitempty"`
	ImageId         string            `json:"imageId,omitempty"`
	DeploymentType  string            `json:"deploymentType,omitempty"`
	OperatorId      string            `json:"operatorId,omitempty"`
	Config          map[string]string `json:"config,omitempty"`
	OutputTopic     string            `json:"outputTopic,omitempty"`
	PersistData     bool              `json:"persistData,omitempty"`
	InputTopics     []InputTopic
	InputSelections []InputSelection `json:"inputSelections,omitempty"`
	Cost            uint             `json:"cost"`
	UpstreamConfig UpstreamConfig `json:"upstream,omitempty"`
	DownstreamConfig DownstreamConfig `json:"downstream,omitempty"`
}

type OperatorRequestConfig struct {
	Config      map[string]string `json:"config,omitempty"`
	InputTopics []InputTopic      `json:"inputTopics,omitempty"`
}

type InputTopic struct {
	Name         string    `json:"name,omitempty"`
	FilterType   string    `json:"filterType,omitempty"`
	FilterValue  string    `json:"filterValue,omitempty"`
	FilterValue2 string    `json:"filterValue2,omitempty"`
	Mappings     []Mapping `json:"mappings,omitempty"`
}

type Mapping struct {
	Dest   string `json:"dest,omitempty"`
	Source string `json:"source,omitempty"`
}

type InputSelection struct {
	InputName         string   `json:"inputName,omitempty"`
	AspectId          string   `json:"aspectId,omitempty"`
	FunctionId        string   `json:"functionId,omitempty"`
	CharacteristicIds []string `json:"characteristicIds,omitempty"`
	SelectableId      string   `json:"selectableId,omitempty"`
}

type Metrics struct {
	Database string `json:"database,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Url      string `json:"url,omitempty"`
	Interval string `json:"interval,omitempty"`
	XmlUrl   string `json:"xmlurl,omitempty"`
}

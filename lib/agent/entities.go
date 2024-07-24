package agent

import (
	"time"

	"github.com/SENERGY-Platform/analytics-fog-lib/lib/control"
)

type AgentInfoMessage struct {
	control.ControlMessage
	Conf Configuration `json:"agent"`
	CurrentOperatorStates []OperatorState
}

type OperatorState struct {
	PipelineID string 
	OperatorID string 
	State string 
	ErrMsg string
	ContainerID string
	Time time.Time `json:"time"`
}

type Configuration struct {
	Id string `json:"id,omitempty"`
}

type Agent struct {
	Id      string    `json:"id,omitempty"`
	Updated time.Time `json:"updated,omitempty"`
	Active  bool      `json:"active,omitempty"`
}

type Ping struct {
	control.ControlMessage
	Updated time.Time `json:"updated,omitempty"`
}

package agent

import (
	"github.com/SENERGY-Platform/analytics-fog-lib/lib/control"
)

type AgentRegisterMessage struct {
	control.ControlMessage
	Conf Configuration `json:"agent"`
}

type AgentPongMessage struct {
	AgentRegisterMessage
}

type Configuration struct {
	Id string `json:"id,omitempty"`
}

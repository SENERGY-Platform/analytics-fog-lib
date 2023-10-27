package master

import (
	"github.com/SENERGY-Platform/analytics-fog-lib/lib/control"
)

type MasterInfoMessage struct {
	control.ControlMessage
	Conf Configuration `json:"master"`
}

type Configuration struct {
	Id string `json:"id,omitempty"`
}

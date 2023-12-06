package agent

import (
	"github.com/SENERGY-Platform/analytics-fog-lib/lib/fog"
)

// Used by Agents to register
const AgentsTopic = fog.FogAnalyticsTopicPrefix + "agents"

func GetStartOperatorAgentTopic(agentID string) string {
	return AgentsTopic + "/" + agentID + "/control/start"
}

func GetStopOperatorAgentTopic(agentID string) string {
	return AgentsTopic + "/" + agentID + "/control/stop"
}
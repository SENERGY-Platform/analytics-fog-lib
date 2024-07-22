package operator

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
}

type Operator struct {
	OperatorIDs
	DeploymentState string                `json:"state"`
	DeploymentError string `json:"deployment_error"`
	AgentId string                `json:"agent_id"`
	ContainerId string `json:"container_id"`
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
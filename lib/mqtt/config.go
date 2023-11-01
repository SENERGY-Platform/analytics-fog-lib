package mqtt

type BrokerConfig struct {
	Host string `json:"broker_host" env_var:"BROKER_HOST"`
	Port string `json:"broker_port" env_var:"BROKER_PORT"`
}

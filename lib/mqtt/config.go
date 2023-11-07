package mqtt

type BrokerConfig struct {
	Host string 
	Port string 
}

type FogBrokerConfig struct {
	Host string `json:"broker_host" env_var:"FOG_BROKER_HOST"`
	Port string `json:"broker_port" env_var:"FOG_BROKER_PORT"`
}

type PlatformBrokerConfig struct {
	Host string `json:"broker_host" env_var:"PLATFORM_BROKER_HOST"`
	Port string `json:"broker_port" env_var:"PLATFORM_BROKER_PORT"`
}


package mqtt

type BrokerConfig struct {
	Host     string
	Port     string
	Username *string
	Password *string
}

type FogBrokerConfig struct {
	Host     string  `json:"broker_host" env_var:"FOG_BROKER_HOST"`
	Port     string  `json:"broker_port" env_var:"FOG_BROKER_PORT"`
	Username *string `json:"username" env_var:"FOG_BROKER_USERNAME"`
	Password *string `json:"password" env_var:"FOG_BROKER_PASSWORD"`
}

type PlatformBrokerConfig struct {
	Host     string  `json:"broker_host" env_var:"PLATFORM_BROKER_HOST"`
	Port     string  `json:"broker_port" env_var:"PLATFORM_BROKER_PORT"`
	Username *string `json:"username" env_var:"AUTH_USERNAME"`
	Password *string `json:"password" env_var:"AUTH_PASSWORD"`
}

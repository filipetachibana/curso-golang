package configuration

import (
	appsettings "../infraestructure/appsettings"
)

// Configuration configuração chaves
type Configuration struct {
	ConnectionString string `json:"ConnectionString"`
	ServerAddress    string `json:"ServerAddress"`
	ServerPort       int    `json:"ServerPort"`
	SubServerName    string `json:"SubServerName"`
}

var _config Configuration

func init() {
	appsettings.SetConfigurationDesenv("C:\\git-go\\inlog-service-alarm-deleteded", &_config)
}

// GetConfiguration configuration
func GetConfiguration() Configuration {
	return _config
}

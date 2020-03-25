package tests

import (
	"encoding/json"

	loggingC "github.com/codemodify/systemkit-logging/contracts"

	helpersConfig "github.com/codemodify/systemkit-config"
)

type AppConfig struct {
	LogLevel int       `json:"logLevel"`
	APIKey   string    `json:"apiKey"`
	Services []Service `json:"services"`
}

type Service struct {
	Listen string `json:"listen"`
}

// Implement interfaces from `github.com/codemodify/systemkit-logging/Config/contracts.go`

func (thisRef *AppConfig) CreateOrReturnInstance() helpersConfig.Config {
	return helpersConfig.LoadConfig(&AppConfig{})
}

func (thisRef *AppConfig) DefaultConfig() helpersConfig.Config {
	return &AppConfig{
		LogLevel: int(loggingC.TypeInfo),
		APIKey:   "",
		Services: []Service{},
	}
}

func (thisRef *AppConfig) String() string {
	bytes, err := json.Marshal(thisRef)
	if err != nil {
		// INFO: in normal app you could log this
		return ""
	}
	return string(bytes)
}

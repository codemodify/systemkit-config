package tests

import (
	"encoding/json"

	config "github.com/codemodify/systemkit-config"
	logging "github.com/codemodify/systemkit-logging"
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

func (thisRef *AppConfig) CreateOrReturnInstance() config.Config {
	return config.LoadConfig(&AppConfig{})
}

func (thisRef *AppConfig) DefaultConfig() config.Config {
	return &AppConfig{
		LogLevel: int(logging.TypeInfo),
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

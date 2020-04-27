package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	files "github.com/codemodify/systemkit-helpers-files"
	reflection "github.com/codemodify/systemkit-helpers-reflection"
	logging "github.com/codemodify/systemkit-logging"
)

var configInstance Config

// LoadConfig - loads a custom config type, takes an existing instance usually created with `&CustomAppConfig{}`
func LoadConfig(config Config) Config {
	configInstance = config.DefaultConfig()

	var configFileName = GetConfigFile()
	if files.FileOrFolderExists(configFileName) {
		file, err := ioutil.ReadFile(configFileName)
		if err != nil {
			logging.Warningf("%s from %", fmt.Sprintf("unable to load config file %s, using default", configFileName), reflection.GetThisFuncName())
		} else {
			err := json.Unmarshal(file, configInstance)
			if err != nil {
				logging.Warningf("%s from %", fmt.Sprintf("unable to load config file %s, using default", configFileName), reflection.GetThisFuncName())

				configInstance = config.DefaultConfig()
			}
		}
	}

	return configInstance
}

// ChangedEventHandler -
type ChangedEventHandler func(config Config)

// OnConfigChanged -
func OnConfigChanged(configChangedEventHandler ChangedEventHandler) {
	files.WatchFile(GetConfigFile(), func() {
		configChangedEventHandler(LoadConfig(configInstance))
	})
}

// SaveConfig -
func SaveConfig(config Config) error {
	return ioutil.WriteFile(GetConfigFile(), []byte(config.String()), 0644)
}

// GetConfigDir -
func GetConfigDir() string {
	directoryName := "." + string(os.PathSeparator) + "config"
	if !files.FileOrFolderExists(directoryName) {
		err := os.MkdirAll(directoryName, 0755)
		if err != nil {
			logging.Fatalf("%s from %", fmt.Sprintf("unable to create directory %s", directoryName), reflection.GetThisFuncName())

			panic(err)
		}
	}
	return directoryName
}

// GetConfigFile -
func GetConfigFile() string {
	return GetConfigDir() + string(os.PathSeparator) + "config.json"
}

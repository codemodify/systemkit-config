package tests

import (
	"fmt"
	"testing"
)

func Test_LoadConfig(t *testing.T) {
	config := getConfig()

	t.Log(config)
	fmt.Println(config)
}

func getConfig() *AppConfig {
	appDefaultConfig := &AppConfig{}
	return appDefaultConfig.CreateOrReturnInstance().(*AppConfig)
}

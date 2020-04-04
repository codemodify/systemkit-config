# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Configuration
[![](https://img.shields.io/github/v/release/codemodify/systemkit-config?style=flat-square)](https://github.com/codemodify/systemkit-config/releases/latest)
![](https://img.shields.io/github/languages/code-size/codemodify/systemkit-config?style=flat-square)
![](https://img.shields.io/github/last-commit/codemodify/systemkit-config?style=flat-square)
[![](https://img.shields.io/badge/license-0--license-brightgreen?style=flat-square)](https://github.com/codemodify/TheFreeLicense)

![](https://img.shields.io/github/workflow/status/codemodify/systemkit-config/qa?style=flat-square)
![](https://img.shields.io/github/issues/codemodify/systemkit-config?style=flat-square)
[![](https://goreportcard.com/badge/github.com/codemodify/systemkit-config?style=flat-square)](https://goreportcard.com/report/github.com/codemodify/systemkit-config)

[![](https://img.shields.io/badge/godoc-reference-brightgreen?style=flat-square)](https://godoc.org/github.com/codemodify/systemkit-config)
![](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)
![](https://img.shields.io/gitter/room/codemodify/systemkit-config?style=flat-square)

![](https://img.shields.io/github/contributors/codemodify/systemkit-config?style=flat-square)
![](https://img.shields.io/github/stars/codemodify/systemkit-config?style=flat-square)
![](https://img.shields.io/github/watchers/codemodify/systemkit-config?style=flat-square)
![](https://img.shields.io/github/forks/codemodify/systemkit-config?style=flat-square)

#### Flexible config framework for your Go app


#### Supported: Linux, Raspberry Pi, FreeBSD, Mac OS, Windows, Solaris

# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Install
```go
go get github.com/codemodify/systemkit-config
```
# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) API

&nbsp;																| &nbsp;
---     															| ---
LoadConfig(`config`) | Loads a custom config type
OnConfigChanged(`configChangedEventHandler`) | On configuration changed
SaveConfig(`config`) | Save configuration
GetConfigDir()  | Get configuration directory 
GetConfigFile() | Get configuration file

# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Usage
```go
package main

import (
	"fmt"
	"testing"

	logging "github.com/codemodify/systemkit-logging"
)

func main(t *testing.T) {
	logging.Init(logging.NewConsoleLogger())

	config := getConfig()

	t.Log(config)
	fmt.Println(config)
}

func getConfig() *AppConfig {
	appDefaultConfig := &AppConfig{}
	return appDefaultConfig.CreateOrReturnInstance().(*AppConfig)
}
```
> - ### for examples see `tests` 
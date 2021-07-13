// Package conf ...
package conf

import (
	"github.com/olebedev/config"
)

//Config ...
var Config = new(config.Config)

// SetCfg format config
func SetCfg(configPath, branch string) {

	cfg, err := config.ParseYamlFile(configPath)

	if err != nil {
		panic(err.Error())
	}

	Config, err = cfg.Get(branch)
	if err != nil {
		panic("config have not " + branch + " branch")
	}
}

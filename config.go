package linqcore

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/SofyanHadiA/linqcore/utils"

	"gopkg.in/yaml.v2"
)

var envarConfigPrefix string

// Configs object
type Configs map[string]interface{}

var configs Configs

// NewConfig create new config class
func NewConfig(envarConfigPrefix string) Configs {
	if len(envarConfigPrefix) > 0 {
		envarConfigPrefix = "LINQ_"
	}

	configs := loadConfig("conf/db.conf")
	appConfig := loadConfig("conf/app.conf")

	utils.MapCopy(configs, appConfig)

	utils.Log.Info("Application configs: ", configs)
	return configs
}

// GetStrConfig get config type string
func (config Configs) GetStrConfig(configKey string) string {
	return config[configKey].(string)
}

// GetIntConfig get config type integer
func (config Configs) GetIntConfig(configKey string) int {
	return config[configKey].(int)
}

func loadConfig(file string) Configs {

	var config Configs

	conf, errFile := ioutil.ReadFile(file)
	if errFile != nil {
		utils.Log.Fatal("error: %v", errFile)
	}

	errYaml := yaml.Unmarshal([]byte(conf), &config)
	if errYaml != nil {
		utils.Log.Fatal("error: %v", errYaml)
	}

	for k := range config {
		envarKey := envarConfigPrefix + strings.ToUpper(strings.Replace(k, ".", "_", -1))
		envarValue := os.Getenv(envarKey)
		if len(envarValue) > 0 {
			config[k] = envarValue
		}
	}

	return config
}

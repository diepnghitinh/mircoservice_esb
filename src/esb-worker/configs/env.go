package configs

import (
	"encoding/json"
	"esb-worker/pkg/envconfig"
	"os"
)

type Env struct {
	config interface{}
}

// getEnv get key environment variable if exist otherwise return defaultValue
func (self *Env) getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

// NewFromEnvironment creates a config object from environment variables
func (self *Env) NewFromEnvironmentFromConfig(_config interface{}) (interface{}, error) {
	self.config = _config
	cnf, err := self.fromEnvironmentConfig()
	if err != nil {
		return nil, err
	}

	log.Info("Successfully loaded config from the environment")

	return cnf, nil
}

func (self *Env) fromEnvironmentConfig() (interface{}, error) {

	var defaultInterface map[string]interface{}
	encode, _ := json.Marshal(*defaultCnf)
	json.Unmarshal(encode, &defaultInterface)

	if err := envconfig.Process("", self.config); err != nil {
		return nil, err
	}

	return self.config, nil
}

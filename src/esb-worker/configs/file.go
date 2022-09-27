package configs

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type File struct {
	config interface{}
}

// NewFromYaml creates a config object from YAML file
func (c *File) NewFromYaml(_config interface{}, cnfPath string) (interface{}, error) {
	c.config = _config
	cnf, err := c.fromFile(cnfPath)
	if err != nil {
		return nil, err
	}

	log.Infof("Successfully loaded config from file %s", cnfPath)

	return cnf, nil
}

// ReadFromFile reads data from a file
func (c *File) ReadFromFile(cnfPath string) ([]byte, error) {
	file, err := os.Open(cnfPath)

	// Config file not found
	if err != nil {
		return nil, fmt.Errorf("Open file error: %s", err)
	}

	// Config file found, let's try to read it
	data := make([]byte, 1000)
	count, err := file.Read(data)
	if err != nil {
		return nil, fmt.Errorf("Read from file error: %s", err)
	}

	return data[:count], nil
}

func (c *File) fromFile(cnfPath string) (interface{}, error) {

	var defaultInterface map[string]interface{}
	encode, _ := yaml.Marshal(*defaultCnf)
	yaml.Unmarshal(encode, &defaultInterface)

	// Assign default
	//*cnf = defaultInterface

	data, err := c.ReadFromFile(cnfPath)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, c.config); err != nil {
		return nil, fmt.Errorf("Unmarshal YAML error: %s", err)
	}

	return c.config, nil
}
package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// this code is generated in https://zhwt.github.io/yaml-to-go/
type Config struct {
	Gateway struct {
		Addr string `yaml:"addr"`
	} `yaml:"gateway"`
	Auth struct {
		Addr string `yaml:"addr"`
	} `yaml:"auth"`
	Course struct {
		Addr string `yaml:"addr"`
	} `yaml:"course"`
	Exam struct {
		Addr string `yaml:"addr"`
	} `yaml:"exam"`
	Database struct {
		Addr   string `yaml:"addr"`
		DbName string `yaml:"db_name"`
	} `yaml:"database"`
	Deployment string `yaml:"deployment"`
}

func LoadConfig() (*Config, error) {
	var conf Config

	yamlFile, err := ioutil.ReadFile("../config/config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return nil, err
	}

	return &conf, nil
}

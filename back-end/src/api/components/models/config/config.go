package config

import (
	// "os"
	"github.com/go-yaml/yaml"
// 	"io/ioutil"
	// "log"
	"fmt"
	"os"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		User string `yaml:"user"`
		Password string `yaml:"pass"`
		Name string `yaml:"name"`
	} `yaml:"database"`
}

func GetConfig() Config {
	var cfg Config
	f, err := os.Open("config.yaml")

	if err != nil {
		fmt.Println(err)
		f.Close()
		return cfg
	}
	decoder := yaml.NewDecoder(f)

	err = decoder.Decode(&cfg)

	if err != nil {
		fmt.Println(err)
		f.Close()
		return cfg
	}

	return cfg
}
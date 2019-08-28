package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Load returns Configuration struct
func Load(path string) (*Configuration, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}
	var cfg = new(Configuration)
	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}
	return cfg, nil
}

// Configuration holds data necessary for configuring application
type Configuration struct {
	APIms *APIms `yaml:"apims,omitempty"`
	Grpc  *Grpc  `yaml:"grpc,omitempty"`
}

// Database holds data necessary for database configuration
type Database struct {
	User     string `yaml:"user,omitempty"`
	Password string `yaml:"password,omitempty"`
	Database string `yaml:"database,omitempty"`
	Addr     string `yaml:"addr,omitempty"`
}

// Server holds data necessary for server configuration
type Server struct {
	Port string `yaml:"port,omitempty"`
}

// JWT holds data necessary for JWT configuration
type JWT struct {
	Secret           string `yaml:"secret,omitempty"`
	Duration         string `yaml:"duration_minutes,omitempty"`
	SigningAlgorithm string `yaml:"signing_algorithm,omitempty"`
}

type APIms struct {
	DB     *Database `yaml:"database,omitempty"`
	JWT    *JWT      `yaml:"jwt,omitempty"`
	Server *Server   `yaml:"server,omitempty"`
}

// Grpc holds data necessary for gRPC configuration
type Grpc struct {
	Port string `yaml:"port,omitempty"`
}

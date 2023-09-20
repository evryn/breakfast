package config

import (
	"log"
	"os"
	"breakfast/types"

	"gopkg.in/yaml.v2"
)

var Main Config

type Config struct {
	MockedVersion string `yaml:"mockedVersion"`

	WebServer struct {
		StaticPath string `yaml:"staticPath"`
		AppPath    string `yaml:"appPath"`
		Interface  string `yaml:"interface"`
		Port       int    `yaml:"port"`
	} `yaml:"webServer"`

	Paths struct {
		StaticDir     string `yaml:"staticDir"`
		LongTemplate  string `yaml:"longTemplate"`
		ShortTemplate string `yaml:"shortTemplate"`
	} `yaml:"paths"`

	Versions         []types.Version         `yaml:"versions"`
	VersionModifiers []types.VersionModifier `yaml:"versionModifiers"`
}

func MustLoadConfig() {
	// Read the YAML file
	yamlData, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Parse the YAML data into the AAA structure
	err = yaml.Unmarshal(yamlData, &Main)
	if err != nil {
		log.Fatalf("Error unmarshaling YAML: %v", err)
	}
}

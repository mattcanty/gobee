package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mattcanty/gobee/pkg/macos"
	"gopkg.in/yaml.v2"
)

type gobeeConfig struct {
	macOS `yaml:"macOS"`
}

// MacOS the configuration for Mac OS
type macOS struct {
	macos.Dock        `yaml:"dock"`
	macos.DateAndTime `yaml:"dateAndTime"`
}

type component interface {
	GetChanges() []string
	ApplyChanges() []string
}

func main() {
	home, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	defaultConfigPath := fmt.Sprintf("%s/.config/gobee/config.yaml", home)
	contents, err := ioutil.ReadFile(defaultConfigPath)
	if err != nil {
		log.Fatalf("Failed to fin Gobee config %s\n", err)
	}

	var cfg gobeeConfig
	err = yaml.Unmarshal(contents, &cfg)
	if err != nil {
		log.Fatalf("Failed to Unmarshal YAMl. %s\n", err)
	}

	// getChanges(cfg.macOS.Dock)
	// getChanges(cfg.macOS.DateAndTime)

	// err = systempreferences.ConfigureDock(cfg.MacOS.Dock)

	// if err != nil {
	// 	log.Fatalf("Failed to ConfigureDock. %s\n", err)
	// }

	// err = systempreferences.ConfigureDateAndTime(cfg.MacOS.DateAndTime)

	// if err != nil {
	// 	log.Fatalf("Failed to ConfigureDateAndTime. %s\n", err)
	// }
}

// func getChanges(c component) {

// }

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/mattcanty/gobee/pkg/config"
	"github.com/mattcanty/gobee/pkg/macos/systempreferences"
)

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

	var cfg config.Config
	err = yaml.Unmarshal(contents, &cfg)
	if err != nil {
		log.Fatalf("Failed to Unmarshal YAMl. %s\n", err)
	}

	err = systempreferences.ConfigureDock(cfg.MacOS.Dock)

	if err != nil {
		log.Fatalf("Failed to ConfigureDock. %s\n", err)
	}

	err = systempreferences.ConfigureDateAndTime(cfg.MacOS.DateAndTime)

	if err != nil {
		log.Fatalf("Failed to ConfigureDateAndTime. %s\n", err)
	}
}

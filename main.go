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
		log.Fatalf("%s\nHave you created a Gobee config file at %s?", err, defaultConfigPath)
	}

	var cfg config.Config
	err = yaml.Unmarshal(contents, &cfg)
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	err = systempreferences.ConfigureDock(cfg.MacOS.Dock)

	if err != nil {
		log.Fatalf("Failed to apply configuration. %s\n", err)
	}
}

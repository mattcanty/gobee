package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/mattcanty/gobee/pkg/config"
	"github.com/mattcanty/gobee/pkg/systempreferences"
)

func main() {
	var fileName string
	flag.StringVar(&fileName, "f", "", "YAML file to parse.")
	flag.Parse()

	if fileName == "" {
		fmt.Println("Please specify your config file using the -f option")
		return
	}

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	defer f.Close()

	var cfg config.Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	systempreferences.ConfigureDock(cfg.Apple.Dock)
}

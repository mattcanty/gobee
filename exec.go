package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Apple struct {
		Dock struct {
			Orientation string `yaml:"orientation"`
		} `yaml:"dock"`
	} `yaml:"apple"`
	Database struct {
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
	} `yaml:"database"`
}

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

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	cmd := exec.Command("defaults", "write", "com.apple.dock", "orientation", "-string", cfg.Apple.Dock.Orientation)
	err = cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

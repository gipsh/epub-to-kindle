package main

import (
	"os"
	"fmt"

        "gopkg.in/yaml.v2"
)

type Config struct {
    SMTP struct {
        Port int `yaml:"port"`
        Host string `yaml:"host"`
        User string `yaml:"user"`
        Pass string `yaml:"pass"`
        Sender string `yaml:"sender"`
    } `yaml:"smtp"`
   WEB struct {
	Port string `yaml:"port"`
    } `yaml:"web"`
}

func processError(err error) {
    fmt.Println(err)
    os.Exit(2)
}


func readFile(cfg *Config) {
    f, err := os.Open("config.yml")
    if err != nil {
        processError(err)
    }
    defer f.Close()

    decoder := yaml.NewDecoder(f)
    err = decoder.Decode(cfg)
    if err != nil {
        processError(err)
    }
} 

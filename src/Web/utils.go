package main

import (
	"fmt"
	"log"
	"encoding/json"
	"os"
	"html/template"
	"net/http"
	"Web/data"
	"strings"
	"errors"
)

// Configuration : config struct model
type Configuration struct {
	Address string
	ReadTimeout int64
	WriteTimeout int64
	Static string
}

var config Configuration
var logger *log.Logger

func p(a ...interface{}) {
	fmt.Println(a)
}

func init() {
	loadConfig()
	file, err := os.OpenFile("Web.log", os.O_CREATE|os.WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal("Cannot open config file", err)
	}
	decoder := json.NewDecode(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("Cannot get config from file", err)
	}
}
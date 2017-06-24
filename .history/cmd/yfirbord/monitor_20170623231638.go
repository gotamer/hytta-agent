package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/lucavallin/yfirbord-grovepi/pkg/connections"
	"github.com/lucavallin/yfirbord-grovepi/pkg/connections/grovepi"
)

func main() {
	// Load sensors
	var config interface{}
	configJSON, err := ioutil.ReadFile("./sensors.config.json")
	if err != nil {
		panic("Sensors configuration file not found. Aborting.")
	}
	if err := json.Unmarshal(configJSON, &config); err != nil {
		panic("Sensors configuration file is not a valid JSON.")
	}

	spew.Dump(config)

	// Init GrovePi on address 0x04
	g, err := connections.grovepi.Init(0x04)
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	// Create readers
	for sensorConf := range config {
		// sensor := sensors.NewSensor()
	}
}

package readers

import (
	"github.com/lucavallin/yfirbord-grovepi/pkg/connections"
	"github.com/lucavallin/yfirbord-grovepi/pkg/sensors"
)

type Reader interface {
	Read() (sensors.Measurement, error)
}

func NewReader(sensor sensors.Sensor) (Reader, error) {
	switch sensor.Mode {
	case "dht":
		return DHT{sensor, conn: connections.GrovePi}, nil
	case "light":

	}

}

package readers

import (
	"github.com/lucavallin/yfirbord-grovepi/pkg/connections"
	"github.com/lucavallin/yfirbord-grovepi/pkg/sensors"
)

type Reader interface {
	Read() (sensors.Measurement, error)
}

func NewReader(sensor sensors.Sensor, conn connections.DHTInput) (Reader, error) {
	conn := new(connections.GrovePi)

	switch sensor.Mode {
	case "dht":
		return DHT{sensor, conn}, nil
		break
	case "light":
		return Light{sensor, conn}, nil
	default:

	}
}
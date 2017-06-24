package sensors

import (
	"time"

	"github.com/lucavallin/yfirbord-grovepi/pkg/grovepi"
)

// Light is structure for DHT sensor
type Light struct {
	Sensor
}

func (o Light) Read() {
	light, err := o.connection.analogRead(grovepi.A2)
	if err != nil {
		panic(err)
	}
	time.Sleep(500 * time.Millisecond)

	return Measurement{
		"light": light,
	}, nil
}

package readers

import (
	"github.com/lucavallin/yfirbord-grovepi/pkg/sensors"
)

type Reader interface {
	Read() (sensors.Measurement, error)
}

func NewReader(sensors.Sensor) (Reader, error) {
	switch expression {
	case condition:

	}
}

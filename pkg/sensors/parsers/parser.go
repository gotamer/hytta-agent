package parsers

// Parser interface for output parsers
type Parser interface {
	ToSI(raw []byte) Measurement
}

// Measurement is a map with the sensor name as index and the measurement as vaue
// e.g. map[temperature] = 32 or map[humidity] = 64
type Measurement map[string]interface{}

// GetParsers returns a map of the register parsers
func GetParsers() map[string]Parser {
	return map[string]Parser{
		"dht":   new(DHT),
		"light": new(Light),
	}
}

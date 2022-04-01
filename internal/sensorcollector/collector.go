package sensorcollector

import (
	"strings"
	"widgetsensor/internal/errors"
)

type Sensor interface {
	Consume(string) error
	ReferenceValid() bool
}

func NewSensor() Sensor {
	reference := &reference{}
	thermometer := newThermometer()

	handlers := map[string]handler{
		"reference":   reference,
		"thermometer": thermometer,
	}

	return &sensor{
		handlers:    handlers,
		reference:   reference,
		thermometer: thermometer,
	}
}

type sensor struct {
	handlers    map[string]handler
	reference   *reference
	thermometer *thermometer
}

func (s *sensor) Consume(line string) error {
	lines := strings.Split(line, " ")
	if len(lines) < 2 {
		return errors.ErrInvalidLine
	}

	if handler, ok := s.handlers[lines[0]]; ok {
		if len(lines) < handler.argLen()+1 {
			return errors.ErrInvalidLine
		}
		err := handler.consume(lines[1:])
		if err != nil {
			return err
		}
	} else {
		if len(lines) < 3 {
			return errors.ErrInvalidLine
		}

		sensorName := lines[1]
		sensorVal := lines[2]

		switch {
		case s.thermometer.isSensor(sensorName):
			s.thermometer.accept(sensorName, sensorVal)
		default:
			return errors.ErrUnknownSensor
		}
	}

	return nil
}

func (s *sensor) ReferenceValid() bool {
	return s.reference.Valid()
}
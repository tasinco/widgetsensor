package sensorcollector

import (
	"strings"
	"widgetsensor/internal/errors"
)

type Sensor interface {
	Consume(string) error
	ReferenceValid() bool
	Output(func(string))
}

func NewSensor() Sensor {
	reference := &reference{}
	thermometer := newThermometer()
	humidity := newHumidity()

	handlers := map[string]handler{
		"reference":   reference,
		"thermometer": thermometer,
		"humidity":    humidity,
	}

	return &sensor{
		handlers:    handlers,
		reference:   reference,
		thermometer: thermometer,
		humidity:    humidity,
	}
}

type sensor struct {
	handlers    map[string]handler
	reference   *reference
	thermometer *thermometer
	humidity    *humidity
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
		case s.humidity.isSensor(sensorName):
			s.humidity.accept(s.reference, sensorName, sensorVal)
		default:
			return errors.ErrUnknownSensor
		}
	}

	return nil
}

func (s *sensor) ReferenceValid() bool {
	return s.reference.Valid()
}

func (h *sensor) Output(out func(string)) {
	h.thermometer.Output(h.reference, out)
	h.humidity.Output(out)
}

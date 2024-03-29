package sensorcollector

import (
	"context"
	"github.com/neilotoole/errgroup"
	"strings"
	"widgetsensor/internal/errors"
)

type Sensor interface {
	Consume(string) error
	ReferenceValid() bool
	Output(func(string))
	Wait() error
}

func NewSensor() Sensor {
	reference := &reference{}
	thermometer := newThermometer()
	humidity := newHumidity()

	errg, _ := errgroup.WithContext(context.Background())

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
		monitors:    make(map[string]sensorMonitor),
		errg:        errg,
	}
}

type sensor struct {
	handlers    map[string]handler
	reference   *reference
	thermometer *thermometer
	humidity    *humidity
	monitors    map[string]sensorMonitor
	errg        *errgroup.Group
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
		sensorName, sensorMonitor, err := handler.consume(lines[1:])
		if err != nil {
			return err
		}
		if sensorMonitor != nil {
			s.monitors[sensorName] = sensorMonitor
		}
	} else {
		if len(lines) < 3 {
			return errors.ErrInvalidLine
		}

		sensorName := lines[1]
		sensorVal := lines[2]

		if sensorMonitor, ok := s.monitors[sensorName]; ok {
			s.errg.Go(func() error {
				return sensorMonitor.accept(s.reference, sensorVal)
			})
		} else {
			return errors.ErrUnknownSensor
		}
	}

	return nil
}

func (s *sensor) Wait() error {
	return s.errg.Wait()
}

func (s *sensor) ReferenceValid() bool {
	return s.reference.Valid()
}

func (h *sensor) Output(out func(string)) {
	h.thermometer.Output(h.reference, out)
	h.humidity.Output(out)
}

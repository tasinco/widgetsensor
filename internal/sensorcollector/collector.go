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

	handlers := map[string]handler{
		"reference": reference,
	}

	return &sensor{
		handlers:  handlers,
		reference: reference,
	}
}

type sensor struct {
	handlers  map[string]handler
	reference *reference
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
	}

	return nil
}

func (s *sensor) ReferenceValid() bool {
	return s.reference.Valid()
}

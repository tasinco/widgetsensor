package sensorcollector

import (
	"bytes"
	"fmt"
	"strconv"
)

type humidity struct {
	monitors map[string]bool
}

func (h humidity) String() string {
	b := bytes.Buffer{}
	for k, v := range h.monitors {
		b.WriteString(fmt.Sprintf("%s: %t; ", k, v))
	}
	return b.String()
}

func newHumidity() *humidity {
	return &humidity{
		monitors: make(map[string]bool),
	}
}

func (h *humidity) isSensor(sensorName string) bool {
	_, ok := h.monitors[sensorName]
	return ok
}

func (h *humidity) Precision(sensorName string) string {
	if res, ok := h.monitors[sensorName]; ok {
		if res {
			return "discarded"
		}
	}
	return "OK"
}

func (h *humidity) accept(reference *reference, sensorName string, val string) error {
	humidity, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return err
	}
	if humidity > reference.humidityMax || humidity < reference.humidityMin {
		h.monitors[sensorName] = true
	}
	return nil
}

func (h humidity) argLen() int {
	return 1
}

func (h *humidity) consume(lines []string) error {
	sensorName := lines[0]
	if _, ok := h.monitors[sensorName]; !ok {
		h.monitors[sensorName] = false
	}
	return nil
}

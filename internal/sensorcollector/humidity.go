package sensorcollector

import (
	"bytes"
	"fmt"
	"strconv"
)

type humiditySensorMonitor struct {
	res bool
}

func (h *humiditySensorMonitor) Precision() string {
	if h.res {
		return "discarded"
	}
	return "OK"
}

func (h *humiditySensorMonitor) accept(reference *reference, val string) error {
	humidity, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return err
	}
	if humidity > reference.humidityMax || humidity < reference.humidityMin {
		h.res = true
	}
	return nil
}

type humidity struct {
	monitors map[string]*humiditySensorMonitor
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
		monitors: make(map[string]*humiditySensorMonitor),
	}
}

func (h humidity) argLen() int {
	return 1
}

func (h *humidity) consume(lines []string) (string, sensorMonitor, error) {
	sensorName := lines[0]
	currentMonitor, ok := h.monitors[sensorName]
	if !ok {
		currentMonitor = &humiditySensorMonitor{}
		h.monitors[sensorName] = currentMonitor
	}
	return sensorName, currentMonitor, nil
}

func (h *humidity) Output(out func(string)) {
	for sensorName, monitor := range h.monitors {
		out(sensorName + ": " + monitor.Precision())
	}
}

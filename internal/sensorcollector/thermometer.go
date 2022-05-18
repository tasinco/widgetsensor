package sensorcollector

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

var (
	roomTempPrec                  = 0.5
	ultraPreciseDeviation float64 = 3
	veryPreciseDeviaiton  float64 = 5
)

type accumulator struct {
	count        int64
	mean         float64
	m2           float64
	stdDeviation float64
}

func (a *accumulator) Precision(base float64) string {
	sensorDeviation := math.Abs(base - a.mean)
	stdDeviation := a.stdDeviation
	if sensorDeviation < roomTempPrec {
		switch {
		case stdDeviation < ultraPreciseDeviation:
			return "ultra precise"
		case stdDeviation < veryPreciseDeviaiton:
			return "very precise"
		}
	}
	return "precise"
}

func (a accumulator) String() string {
	return fmt.Sprintf("count %d mean %f stddev %f", a.count, a.mean, a.stdDeviation)
}

func (a *accumulator) accumulate(val float64) {
	// compute standard deviation on the fly
	// from: https://en.wikipedia.org/wiki/Algorithms_for_calculating_variance
	a.count++
	delta := val - a.mean
	a.mean += delta / float64(a.count)
	a.m2 += delta * (val - a.mean)
	a.stdDeviation = math.Sqrt(a.m2 / float64(a.count))
}

type thermometerSensorMonitor struct {
	accumulator *accumulator
}

func (t *thermometerSensorMonitor) accept(_ *reference, val string) error {
	thermometerValue, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return err
	}
	t.accumulator.accumulate(thermometerValue)
	return nil
}

func (t *thermometerSensorMonitor) Precision(base float64) string {
	return t.accumulator.Precision(base)
}

type thermometer struct {
	monitors map[string]*thermometerSensorMonitor
}

func (t thermometer) String() string {
	b := bytes.Buffer{}
	for k, v := range t.monitors {
		b.WriteString(fmt.Sprintf("%s: [%v]; ", k, v))
	}
	return b.String()
}

func newThermometer() *thermometer {
	return &thermometer{
		monitors: make(map[string]*thermometerSensorMonitor),
	}
}

func (t thermometer) argLen() int {
	return 1
}

func (t *thermometer) consume(lines []string) (string, sensorMonitor, error) {
	sensorName := lines[0]
	curentMonitor, ok := t.monitors[sensorName]
	if !ok {
		curentMonitor = &thermometerSensorMonitor{accumulator: &accumulator{}}
		t.monitors[sensorName] = curentMonitor
	}
	return sensorName, curentMonitor, nil
}

func (t *thermometer) Output(reference *reference, out func(string)) {
	for sensorName, monitor := range t.monitors {
		out(sensorName + ": " + monitor.Precision(reference.thermometer))
	}
}

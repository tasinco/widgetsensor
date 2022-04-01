package sensorcollector

import (
	"fmt"
	"strconv"
	"widgetsensor/internal/errors"
)

type reference struct {
	valid       bool
	thermometer float64
	humidity    float64
	humidityMax float64
	humidityMin float64
}

func (r reference) String() string {
	return fmt.Sprintf("thermometer %f humidity %f humidity max %f humidity min %f",
		r.thermometer, r.humidity, r.humidityMax, r.humidityMin)
}

func (r reference) argLen() int {
	return 2
}

func (r *reference) consume(lines []string) (string, sensorMonitor, error) {
	newThermometer, err := strconv.ParseFloat(lines[0], 64)
	if err != nil {
		return "", nil, errors.ErrInvalidFloat
	}
	newHumidity, err := strconv.ParseFloat(lines[1], 64)
	if err != nil {
		return "", nil, errors.ErrInvalidFloat
	}

	r.thermometer = newThermometer
	r.humidity = newHumidity

	onePercentHumidity := r.humidity * 0.01
	r.humidityMax = r.humidity + onePercentHumidity
	r.humidityMin = r.humidity - onePercentHumidity

	r.valid = true

	return "", nil, nil
}

func (r *reference) Valid() bool {
	return r.valid
}

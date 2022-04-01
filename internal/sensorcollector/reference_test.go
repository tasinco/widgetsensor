package sensorcollector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReference(t *testing.T) {
	tests := []struct {
		name           string
		lines          []string
		expThermometer float64
		expHumidity    float64
		expHumidityMax float64
		expHumidityMin float64
	}{
		{
			name:           "reference",
			lines:          []string{"10", "20"},
			expThermometer: 10,
			expHumidity:    20,
			expHumidityMax: 20.2,
			expHumidityMin: 19.8,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reference := &reference{}
			assert.False(t, reference.Valid())
			reference.consume(test.lines)
			assert.Equal(t, test.expThermometer, reference.thermometer)
			assert.Equal(t, test.expHumidity, reference.humidity)
			assert.Equal(t, test.expHumidityMin, reference.humidityMin)
			assert.Equal(t, test.expHumidityMax, reference.humidityMax)
			assert.True(t, reference.Valid())
		})
	}
}

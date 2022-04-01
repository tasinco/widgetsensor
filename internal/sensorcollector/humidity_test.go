package sensorcollector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHumidity(t *testing.T) {
	humidity := newHumidity()
	assert.False(t, humidity.isSensor("sensor"))
	humidity.consume([]string{"sensor"})
	assert.True(t, humidity.isSensor("sensor"))
}

func TestHumidityDiscard(t *testing.T) {
	tests := []struct {
		name        string
		reference   []string
		humidityVal string
		expRes      string
	}{
		{
			name:        "humidity ok",
			reference:   []string{"10", "10"},
			humidityVal: "10",
			expRes:      "OK",
		},
		{
			name:        "humidity discard max",
			reference:   []string{"10", "10"},
			humidityVal: "11",
			expRes:      "discarded",
		},
		{
			name:        "humidity discard min",
			reference:   []string{"10", "10"},
			humidityVal: "9",
			expRes:      "discarded",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reference := &reference{}
			reference.consume(test.reference)
			humidity := newHumidity()
			humidity.accept(reference, "sensor", test.humidityVal)
			assert.Equal(t, test.expRes, humidity.Precision("sensor"))
		})
	}
}

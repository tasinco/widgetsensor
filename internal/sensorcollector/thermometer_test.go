package sensorcollector

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestThermometer(t *testing.T) {
	thermometer := newThermometer()
	assert.False(t, thermometer.isSensor("sensor"))
	thermometer.consume([]string{"sensor"})
	assert.True(t, thermometer.isSensor("sensor"))
}

func TestAccumulator(t *testing.T) {
	tests := []struct {
		name        string
		floats      []float64
		expMean     float64
		expDev      float64
		expBranding string
		baseTemp    float64
	}{
		{
			name: "ultra precise",
			floats: []float64{
				69.5,
				70.1,
				71.3,
				71.5,
				69.8,
			},
			expMean:     70.44,
			expDev:      0.808949936646268,
			expBranding: "ultra precise",
			baseTemp:    70,
		},
		{
			name: "very precise",
			floats: []float64{
				66,
				69.1,
				71.9,
				75,
			},
			expMean:     70.5,
			expDev:      3.332416540590328,
			expBranding: "very precise",
			baseTemp:    70.2,
		},
		{
			name: "precise",
			floats: []float64{
				69.5,
				70.1,
				71.3,
				71.5,
				10,
			},
			expMean:     58.48,
			expDev:      24.251383465691188,
			expBranding: "precise",
			baseTemp:    70,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			accumulator := &accumulator{}
			for _, flv := range test.floats {
				accumulator.accumulate(flv)
			}
			assert.Equal(t, test.expBranding, accumulator.Precision(test.baseTemp))
			assert.Equal(t, test.expDev, accumulator.stdDeviation)
			assert.Equal(t, test.expMean, accumulator.mean)
			stdDev := calculateStdDeviation(test.floats)
			delta := math.Abs(stdDev - accumulator.stdDeviation)
			assert.True(t, delta < 0.0000001)
		})
	}
}

func calculateStdDeviation(fls []float64) float64 {
	var sum float64
	for _, fl := range fls {
		sum += fl
	}
	mean := sum / float64(len(fls))

	var sd float64
	for j := 0; j < len(fls); j++ {
		sd += math.Pow(fls[j]-mean, 2)
	}
	return math.Sqrt(sd / float64(len(fls)))
}

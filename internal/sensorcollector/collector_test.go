package sensorcollector

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"widgetsensor/internal/errors"
)

func TestCollectorReference(t *testing.T) {
	tests := []struct {
		name   string
		lines  []string
		expRef bool
		expErr error
	}{
		{
			name: "reference",
			lines: []string{
				"reference 1 2",
			},
			expRef: true,
			expErr: nil,
		},
		{
			name:   "no reference",
			lines:  []string{},
			expRef: false,
			expErr: nil,
		},
		{
			name: "bad reference",
			lines: []string{
				"reference 1",
			},
			expRef: false,
			expErr: errors.ErrInvalidLine,
		},
		{
			name: "bad reference invalid float",
			lines: []string{
				"reference x y",
			},
			expRef: false,
			expErr: errors.ErrInvalidFloat,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sensorCollector := NewSensor()
			assert.False(t, sensorCollector.(*sensor).reference.Valid())
			for _, line := range test.lines {
				err := sensorCollector.Consume(line)
				if test.expErr == nil {
					assert.NoError(t, err)
				} else {
					assert.Equal(t, test.expErr, err)
				}
			}
			assert.Equal(t, test.expRef, sensorCollector.(*sensor).reference.Valid())
		})
	}
}

func TestCollectorThermometer(t *testing.T) {
	tests := []struct {
		name   string
		lines  []string
		expErr []error
	}{
		{
			name: "thermometer",
			lines: []string{
				"thermometer temp-1",
				"2007-04-05T22:00 temp-1 72.4",
			},
			expErr: []error{nil, nil},
		},
		{
			name: "thermometer invalid",
			lines: []string{
				"thermometer",
			},
			expErr: []error{errors.ErrInvalidLine},
		},
		{
			name: "thermometer missing declaration",
			lines: []string{
				"2007-04-05T22:00 temp-1 72.4",
			},
			expErr: []error{errors.ErrUnknownSensor},
		},
		{
			name: "thermometer invalid temp",
			lines: []string{
				"thermometer temp-1",
				"2007-04-05T22:00 temp-1",
			},
			expErr: []error{nil, errors.ErrInvalidLine},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sensorCollector := NewSensor()
			assert.False(t, sensorCollector.(*sensor).reference.Valid())
			err := sensorCollector.Consume("reference 1 2")
			assert.NoError(t, err)
			for cnt, line := range test.lines {
				err := sensorCollector.Consume(line)
				if test.expErr[cnt] == nil {
					assert.NoError(t, err)
				} else {
					assert.Equal(t, test.expErr[cnt], err)
				}
			}
		})
	}
}

func TestCollectorHumidity(t *testing.T) {
	tests := []struct {
		name   string
		lines  []string
		expErr []error
	}{
		{
			name: "humidity",
			lines: []string{
				"humidity hum-1",
				"2007-04-05T22:00 hum-1 72.4",
			},
			expErr: []error{nil, nil},
		},
		{
			name: "humidity invalid",
			lines: []string{
				"humidity",
			},
			expErr: []error{errors.ErrInvalidLine},
		},
		{
			name: "humidity missing declaration",
			lines: []string{
				"2007-04-05T22:00 hum-1 72.4",
			},
			expErr: []error{errors.ErrUnknownSensor},
		},
		{
			name: "humidity invalid temp",
			lines: []string{
				"humidity hum-1",
				"2007-04-05T22:00 hum-1",
			},
			expErr: []error{nil, errors.ErrInvalidLine},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sensorCollector := NewSensor()
			assert.False(t, sensorCollector.(*sensor).reference.Valid())
			err := sensorCollector.Consume("reference 1 2")
			assert.NoError(t, err)
			for cnt, line := range test.lines {
				err := sensorCollector.Consume(line)
				if test.expErr[cnt] == nil {
					assert.NoError(t, err)
				} else {
					assert.Equal(t, test.expErr[cnt], err)
				}
			}
			err = sensorCollector.Wait()
			assert.NoError(t, err)
		})
	}
}

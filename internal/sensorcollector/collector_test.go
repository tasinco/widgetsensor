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

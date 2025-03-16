package entity

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ErrorInvalidMetricsValue = errors.New("invalid metrics value")
var ErrorInvalidMetricsType = errors.New("invalid metrics type")

func TestNewMetrics(t *testing.T) {
	tests := []struct {
		name   string
		mName  string
		mValue string
		mType  string
		want   error
	}{
		{
			name:   "Correct gauge metrics",
			mName:  "someValue",
			mValue: "3.14",
			mType:  "gauge",
			want:   nil,
		},
		{
			name:   "Correct counter metrics",
			mName:  "someValue",
			mValue: "3",
			mType:  "counter",
			want:   nil,
		},
		{
			name:   "Invalid metrics type",
			mName:  "someValue",
			mValue: "3.14",
			mType:  "int",
			want:   ErrorInvalidMetricsType,
		},
		{
			name:   "Invalid counter metrics value",
			mName:  "someValue",
			mValue: "3.1415",
			mType:  "counter",
			want:   ErrorInvalidMetricsValue,
		},
		{
			name:   "Invalid gauge metrics value",
			mName:  "someValue",
			mValue: "test",
			mType:  "gauge",
			want:   ErrorInvalidMetricsValue,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := NewMetrics(test.mName, test.mValue, test.mType)
			if err != nil {
				assert.Equal(t, err.Error(), test.want.Error())
			} else {
				assert.Equal(t, test.want, nil)
			}
		})
	}
}

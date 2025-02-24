package metrix

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGaugeValue(t *testing.T) {
	type want struct {
		create string
		isNill bool
		update string
		err    error
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "Correct gauge value",
			want: want{
				create: "3.0",
				isNill: false,
				update: "3.14",
				err:    nil,
			},
		},
		{
			name: "Invalid gauge value",
			want: want{
				create: "asd",
				isNill: true,
				update: "asd",
				err:    errors.New("invalid"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := NewGaugeValue(test.want.create)
			assert.Equal(t, v == nil, test.want.isNill)
			if v != nil {
				err := v.Update(test.want.update)
				if err != nil {
					assert.Equal(t, err.Error(), test.want.err.Error())
				} else {
					assert.Equal(t, test.want.err, nil)
				}
			}
		})
	}
}

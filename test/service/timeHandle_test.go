package service_test

import (
	"QM-table/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSwitchTimeToLunarCalendar(t *testing.T) {
	type args struct {
		name  string
		times time.Time
		want  []string
		isErr bool
	}
	test := []args{
		{
			name:  "test success",
			times: time.Unix(1694504843, 0),
			want:  []string{"癸卯", "辛酉", "癸酉", "庚申"},
			isErr: false,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := service.SwitchTimeToLunarCalendar(tt.times)
			assert.Equal(t, got, tt.want)
		})
	}

}

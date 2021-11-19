package settings_test

import (
	"proyectos/src/api/config/settings"
	"reflect"
	"testing"
)

func TestGetData(t *testing.T) {
	tests := []struct {
		name string
		want *settings.Data
	}{
		// TODO: Add test cases.
		{name: "first test", want: &settings.Data{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := settings.GetData(); !(reflect.TypeOf(got) == reflect.TypeOf(tt.want)) {
				t.Errorf("GetData() = %v, want %v", got, tt.want)
			}
		})
	}
}

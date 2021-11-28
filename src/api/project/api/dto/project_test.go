package dto

import (
	"proyectos/src/api/project/domain/model"
	"reflect"
	"testing"
)

func TestThing_ToModel(t *testing.T) {
	tests := []struct {
		name string
		t    *Thing
		want *model.Thing
	}{
		// TODO: Add test cases.
		{
			name: "first test",
			t:    &Thing{ID: 1, Name: "algo"},
			want: &model.Thing{ID: 1, Name: "Algo", IsDeleted: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.ToModel(); !(reflect.TypeOf(got) == reflect.TypeOf(tt.want)) {
				t.Errorf("Thing.ToModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThing_FromModel(t *testing.T) {
	type args struct {
		dm *model.Thing
	}
	tests := []struct {
		name string
		t    *Thing
		args args
	}{
		// TODO: Add test cases.
		{
			name: "first test",
			t:    &Thing{ID: 1, Name: "algo"},
			args: args{dm: &model.Thing{ID: 1, Name: "Algo", IsDeleted: false}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.t.FromModel(tt.args.dm)
		})
	}
}

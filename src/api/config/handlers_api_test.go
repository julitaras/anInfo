package config_test

import (
	"proyectos/src/api/config"
	"proyectos/src/api/thing/domain"
	"proyectos/src/api/thing/repository"
	"proyectos/src/api/thing/service"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
)

func TestSRV_AddHandlers(t *testing.T) {
	tests := []struct {
		name string
		r    *config.SRV
		want *config.SRV
	}{
		// TODO: Add test cases.
		{name: "first test", r: config.NewServer(gin.Default()), want: &config.SRV{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.AddHandlers(); !(reflect.TypeOf(got) == reflect.TypeOf(tt.want)) {
				t.Errorf("SRV.AddHandlers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addThingHandler(t *testing.T) {

	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dr := &repository.ThingRepository{DB: db}
	ds := service.NewThingService(dr)

	type args struct {
		r  *config.SRV
		dr domain.Repository
		ds domain.Service
	}
	tests := []struct {
		name string
		args args
		want *config.SRV
	}{
		// TODO: Add test cases.
		{
			name: "first test",
			args: args{
				r:  config.NewServer(gin.Default()),
				dr: dr,
				ds: ds,
			},
			want: &config.SRV{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := config.AddThingHandler(tt.args.r, tt.args.dr, tt.args.ds); !(reflect.TypeOf(got) == reflect.TypeOf(tt.want)) {
				t.Errorf("addThingHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

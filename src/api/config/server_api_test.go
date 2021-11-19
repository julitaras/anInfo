package config

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNewServer(t *testing.T) {
	type args struct {
		g *gin.Engine
	}
	tests := []struct {
		name string
		args args
		want *SRV
	}{
		// TODO: Add test cases.
		{name: "first test", args: args{g: gin.New()}, want: &SRV{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServer(tt.args.g); !(reflect.TypeOf(got) == reflect.TypeOf(tt.want)) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

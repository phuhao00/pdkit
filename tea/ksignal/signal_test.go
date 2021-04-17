package ksignal

import (
	"os"
	"reflect"
	"syscall"
	"testing"
)

func TestHandleSignal(t *testing.T) {
	type args struct {
		c chan os.Signal
		f map[syscall.Signal]func() error
	}
	tests := []struct {
		name string
		args args
		want func() error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HandleSignal(tt.args.c, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandleSignal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitSignal(t *testing.T) {
	tests := []struct {
		name string
		want chan os.Signal
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitSignal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitSignal() = %v, want %v", got, tt.want)
			}
		})
	}
}

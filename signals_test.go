package signals_test

import (
	"syscall"
	"testing"

	"czechia.dev/signals"
)

func TestInterrupt(t *testing.T) {
	type args struct {
		stopFunc func() error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "sigint",
			args: args{
				stopFunc: func() error { return nil },
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			signals.Signals <- syscall.SIGINT
			if err := signals.Interrupt(tt.args.stopFunc); (err != nil) != tt.wantErr {
				t.Errorf("Interrupt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package utils

import (
	"testing"
	"time"
)

func TestFormatTime(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{time.Now().Add(-time.Hour * 24 * 365 * 2)}, "2年0天"},
		{"test2", args{time.Now().Add(-time.Hour * 24 * 365)}, "1年0天"},
		{"test3", args{time.Now().Add(-time.Hour * 24 * 30)}, "30天"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatTime(tt.args.t); got != tt.want {
				t.Errorf("FormatTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

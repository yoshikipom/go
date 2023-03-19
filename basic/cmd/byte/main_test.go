package main

import "testing"

func TestByteToString(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"success", args{b: []byte{0x41}}, "A"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteToString(tt.args.b); got != tt.want {
				t.Errorf("ByteToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

package utils

import "testing"

func TestMD5(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{s: "hello"},
			want: "5d41402abc4b2a76b9719d911017c592",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := MD5(tt.args.s); got != tt.want {
				t.Errorf("MD5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{s: "hello"},
			want: "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA1(tt.args.s); got != tt.want {
				t.Errorf("SHA1() = %v, want %v", got, tt.want)
			}
		})
	}
}

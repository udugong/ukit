package base62

import (
	"reflect"
	"testing"
)

var defaultEntity = NewEntity()

func TestEntity_Encode(t *testing.T) {
	tests := []struct {
		name string
		num  int64
		want string
	}{
		{
			name: "value_of_zero",
			num:  0,
			want: "0",
		},
		{
			name: "normal",
			num:  871246,
			want: "3eeM",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := defaultEntity.Encode(tt.num); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_Decode(t *testing.T) {
	tests := []struct {
		name    string
		encoded string
		want    int64
		wantErr bool
	}{
		{
			name:    "value_of_zero",
			encoded: "0",
			want:    0,
		},
		{
			name:    "normal",
			encoded: "3eeM",
			want:    871246,
		},
		{
			name:    "not_in_ASCIIChars",
			encoded: "=",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := defaultEntity.Decode(tt.encoded)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithASCIIChars(t *testing.T) {
	tests := []struct {
		name  string
		chars string
		want  string
	}{
		{
			name:  "normal",
			chars: "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz~!@#$%^&*()",
			want:  "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz~!@#$%^&*()",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewEntity(WithASCIIChars(tt.chars))
			if got := e.ASCIIChars; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithASCIIChars() = %v, want %v", got, tt.want)
			}
		})
	}
}

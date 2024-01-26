package baseconv

import (
	"math/big"
	"reflect"
	"testing"
)

func TestBaseN_NumToBaseString(t *testing.T) {
	b := NewBaseN(62)
	tests := []struct {
		name string
		num  int64
		want string
	}{
		{
			name: "normal_0",
			num:  0,
			want: "0",
		},
		{
			name: "normal_10",
			num:  10,
			want: "a",
		},
		{
			name: "normal_4592",
			num:  4592,
			want: "1c4",
		},
		{
			// 负数
			name: "normal_-4592",
			num:  -4592,
			want: "-1c4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.NumToBaseString(tt.num); got != tt.want {
				t.Errorf("NumToBaseString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseN_BaseStringToNum(t *testing.T) {
	b := NewBaseN(62)
	tests := []struct {
		name  string
		str   string
		want  int64
		want1 bool
	}{
		{
			name:  "normal_0",
			str:   "0",
			want:  0,
			want1: true,
		},
		{
			name:  "normal_1c4",
			str:   "1c4",
			want:  4592,
			want1: true,
		},
		{
			// 负数
			name:  "normal_-1c4",
			str:   "-1c4",
			want:  -4592,
			want1: true,
		},
		{
			// 有字符不在 chars 中
			name: "not_in_chars_#12",
			str:  "#12",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := b.BaseStringToNum(tt.str)
			if got != tt.want {
				t.Errorf("BaseStringToNum() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("BaseStringToNum() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewBaseN(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		want      *BaseN
		wantPanic bool
	}{
		{
			name:      "normal",
			n:         62,
			want:      &BaseN{n: 62, bInt: big.NewInt(62)},
			wantPanic: false,
		},
		{
			name:      "invalid_base",
			n:         0,
			wantPanic: true,
		},
		{
			name:      "invalid_base",
			n:         -1,
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if tt.wantPanic {
					if r == nil {
						t.Error("want panic but got nil")
					}
				} else {
					if r != nil {
						t.Error("don't want panic but get")
					}
				}
			}()
			if got := NewBaseN(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBaseN() = %v, want %v", got, tt.want)
			}
		})
	}
}

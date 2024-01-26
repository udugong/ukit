package baseconv

import (
	"reflect"
	"testing"
)

func TestCustomBaseN_NumToBaseNString(t *testing.T) {
	c := NewCustomBaseN(2)
	tests := []struct {
		name  string
		baseN int
		chars string
		num   int64
		want  string
	}{
		{
			name: "normal_base2_10",
			num:  10,
			want: "1010",
		},
		{
			name: "less_than_0",
			num:  -1,
			want: "",
		},
		{
			name: "normal_0",
			num:  0,
			want: "0",
		},
		{
			name:  "normal_base62_4592",
			baseN: 62,
			num:   4592,
			want:  "1c4",
		},
		{
			name:  "normal_base62_another_chars_4592",
			baseN: 62,
			chars: "0d13r5qtTD2W9abcOevQfRghPjl6k7mnpUVuSwZxzABCiEF8GHIJs4KLMoNyXY",
			num:   4592,
			want:  "d9r",
		},
	}
	for _, tt := range tests {
		if tt.baseN != 0 {
			c.n = tt.baseN
		}
		if tt.chars != "" {
			c.chars = tt.chars
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := c.NumToBaseNString(tt.num); got != tt.want {
				t.Errorf("NumToBaseNString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomBaseN_BaseNStringToNum(t *testing.T) {
	c := NewCustomBaseN(2)
	tests := []struct {
		name  string
		baseN int
		chars string
		str   string
		want  int64
		want1 bool
	}{
		{
			name:  "normal_base2_1010",
			baseN: 2,
			str:   "1010",
			want:  10,
			want1: true,
		},
		{
			// 2 进制不在 chars 范围中
			name:  "base2_not_in_chars_a",
			baseN: 2,
			str:   "a",
		},
		{
			name:  "normal_base62_1c4",
			baseN: 62,
			str:   "1c4",
			want:  4592,
			want1: true,
		},
		{
			name:  "base62_not_in_chars_-123",
			baseN: 62,
			str:   "-123",
		},
		{
			name:  "normal_base63_1-",
			baseN: 63,
			chars: chars + "-",
			str:   "1-",
			want:  125,
			want1: true,
		},
		{
			name:  "normal_base62_another_chars_4592",
			baseN: 62,
			chars: "0d13r5qtTD2W9abcOevQfRghPjl6k7mnpUVuSwZxzABCiEF8GHIJs4KLMoNyXY",
			str:   "d9r",
			want:  4592,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.baseN != 0 {
				c.n = tt.baseN
			}
			if tt.chars != "" {
				c.chars = tt.chars
			}
			got, got1 := c.BaseNStringToNum(tt.str)
			if got != tt.want {
				t.Errorf("BaseNStringToNum() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("BaseNStringToNum() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewCustomBaseN(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		want      *CustomBaseN
		wantPanic bool
	}{
		{
			name:      "normal",
			n:         62,
			want:      &CustomBaseN{n: 62, chars: chars},
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
			if got := NewCustomBaseN(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCustomBaseN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithSetChars(t *testing.T) {

	tests := []struct {
		name  string
		chars string
		want  string
	}{
		{
			name:  "normal",
			chars: "0d13r5qtTD2W9abcOevQfRghPjl6k7mnpUVuSwZxzABCiEF8GHIJs4KLMoNyXY",
			want:  "0d13r5qtTD2W9abcOevQfRghPjl6k7mnpUVuSwZxzABCiEF8GHIJs4KLMoNyXY",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCustomBaseN(62, WithSetChars(tt.chars)).chars; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithSetChars() = %v, want %v", got, tt.want)
			}
		})
	}
}

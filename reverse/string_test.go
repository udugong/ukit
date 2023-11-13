package reverse

import "testing"

func TestString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal",
			s:    "abc123=0+]",
			want: "]+0=321cba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.s); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRuneString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal",
			s:    "中文啊",
			want: "啊文中",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RuneString(tt.s); got != tt.want {
				t.Errorf("RuneString() = %v, want %v", got, tt.want)
			}
		})
	}
}

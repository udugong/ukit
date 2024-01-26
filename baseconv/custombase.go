package baseconv

import (
	"strings"

	"github.com/udugong/ukit/reverse"
)

const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// CustomBaseN 定义自定义的十进制与 N 进制的转换.
type CustomBaseN struct {
	n     int
	chars string
}

// NewCustomBaseN 创建自定义进制转换.
// 如果 n < 2 或 n > len(CustomBaseN.chars) 则会 panic.
func NewCustomBaseN(n int, opts ...Option) *CustomBaseN {
	res := &CustomBaseN{
		n:     n,
		chars: chars,
	}
	for _, opt := range opts {
		opt.apply(res)
	}
	if n < 2 || n > len(res.chars) {
		panic("invalid base")
	}
	return res
}

type Option interface {
	apply(*CustomBaseN)
}

type optionFunc func(*CustomBaseN)

func (f optionFunc) apply(n *CustomBaseN) {
	f(n)
}

// WithSetChars 设置字符串
func WithSetChars(chars string) Option {
	return optionFunc(func(n *CustomBaseN) {
		n.chars = chars
	})
}

// NumToBaseNString 十进制转换为 N 进制字符串.
// 如果 num < 0 则返回 "".
func (c CustomBaseN) NumToBaseNString(num int64) string {
	if num < 0 {
		return ""
	}
	if num == 0 {
		return string(c.chars[0])
	}

	var src []byte
	base := int64(c.n)
	for num > 0 {
		remainder := num % base
		src = append(src, c.chars[remainder])
		num = num / base
	}

	// 反转
	return string(reverse.Slice(src))
}

// BaseNStringToNum N 进制字符串转十进制.
// 如果有字符不在 chars 中则返回 bool = false.
func (c CustomBaseN) BaseNStringToNum(str string) (int64, bool) {
	base := int64(c.n)
	chars := c.chars[:c.n]
	src := []byte(str)

	var num int64
	for _, b := range src {
		index := strings.IndexByte(chars, b)
		if index == -1 {
			return 0, false
		}
		num = num*base + int64(index)
	}

	return num, true
}

package base62

import (
	"fmt"
	"strings"

	"github.com/udugong/ukit/option"
	"github.com/udugong/ukit/reverse"
)

type Entity struct {
	ASCIIChars string
}

func NewEntity(opts ...option.Option[Entity]) *Entity {
	dOpts := Entity{
		ASCIIChars: "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
	}
	for _, opt := range opts {
		opt.Apply(&dOpts)
	}
	return &dOpts
}

func WithASCIIChars(chars string) option.Option[Entity] {
	return option.NewFuncOption[Entity](func(e *Entity) {
		e.ASCIIChars = chars
	})
}

func (e *Entity) Encode(num int64) string {
	if num == 0 {
		return string(e.ASCIIChars[0])
	}

	var builder strings.Builder
	base := int64(len(e.ASCIIChars))
	for num > 0 {
		remainder := num % base
		builder.WriteByte(e.ASCIIChars[remainder])
		num = num / base
	}
	encoded := builder.String()

	// 反转字符串
	return reverse.String(encoded)
}

func (e *Entity) Decode(encoded string) (int64, error) {
	base := int64(len(e.ASCIIChars))
	var num int64
	for _, char := range encoded {
		index := strings.IndexByte(e.ASCIIChars, byte(char))
		if index == -1 {
			return 0, fmt.Errorf("invalid character: %c", char)
		}
		num = num*base + int64(index)
	}

	return num, nil
}

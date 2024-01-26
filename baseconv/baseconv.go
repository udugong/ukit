package baseconv

import "math/big"

// BaseN 定义十进制与 N 进制转换
type BaseN struct {
	n    int
	bInt *big.Int
}

// NewBaseN 创建十进制与 N 进制转换器
func NewBaseN(n int) *BaseN {
	if n < 2 || n > 62 {
		panic("invalid base")
	}
	return &BaseN{
		n:    n,
		bInt: big.NewInt(int64(n)),
	}
}

// NumToBaseString 十进制转换成 N 进制字符串
func (b *BaseN) NumToBaseString(num int64) string {
	return b.bInt.SetInt64(num).Text(b.n)
}

// BaseStringToNum N 进制字符串转换成十进制
func (b *BaseN) BaseStringToNum(str string) (int64, bool) {
	tmp, ok := b.bInt.SetString(str, b.n)
	if !ok {
		return 0, ok
	}
	return tmp.Int64(), ok
}

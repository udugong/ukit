package baseconv

type Converter interface {
	NumToBaseNString(num int64) string
	BaseNStringToNum(str string) (int64, bool)
}

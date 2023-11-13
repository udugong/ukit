package option

import (
	"testing"
)

func TestNewFuncOption(t *testing.T) {
	o := NewTestOptions()
	if o.Name != "" {
		t.Fail()
	}
}

func TestWithName(t *testing.T) {
	o := NewTestOptions(WithName("foo"))
	if o.Name != "foo" {
		t.Error("设置失败")
	}
}

type TestOptions struct {
	Name string
}

func NewTestOptions(opts ...Option[TestOptions]) TestOptions {
	dOpts := TestOptions{}
	for _, opt := range opts {
		opt.Apply(&dOpts)
	}
	return dOpts
}

func WithName(name string) Option[TestOptions] {
	return NewFuncOption[TestOptions](func(t *TestOptions) {
		t.Name = name
	})
}

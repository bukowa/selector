package selector_test

import (
	. "github.com/bukowa/selector"
	"testing"
)

func TestBaseRegistry_Register(t *testing.T) {
	reg := &BaseRegistry{}
	reg.Register("1", NewBaseSelector("2", func(i interface{}) []Score {
		return nil
	}))
	if len(reg.Scored("")) > 0 {
		t.Error()
	}
}

var benchReg = BaseRegistry{
	"1": {StringSuffix.New("9", "1")},
	"2": {StringPrefix.New("9", "1")},
	"3": {StringEqual.New("9", "1")},
	"4": {StringEqualFold.New("9", "1")},
	"5": {StringContains.New("9", "1")},
	"6": {StringSuffix.New("9", "1")},
}

func BenchmarkBaseRegistry_Scored(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchReg.Scored("yes")
	}
}
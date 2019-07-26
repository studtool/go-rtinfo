package rtinfo

import (
	"testing"

	"github.com/franela/goblin"
)

func testFunc() {}

type testStruct struct{}

func (testStruct) testMethod() {}

func TestGetFuncInfo(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("GetFuncInfo", func() {
		g.It("should return function info", func() {
			info := GetFuncInfo(testFunc)
			g.Assert(info.NameWithPkg).Equal(
				"github.com/studtool/go-rtinfo/pkg/rtinfo.testFunc",
			)
		})
		g.It("should return method info", func() {
			v := testStruct{}
			info := GetFuncInfo(v.testMethod)
			g.Assert(info.NameWithPkg).Equal(
				"github.com/studtool/go-rtinfo/pkg/rtinfo.testStruct.testMethod-fm",
			)
		})
	})
}

func BenchmarkGetFuncInfo_Func(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetFuncInfo(testFunc)
	}
}

func BenchmarkGetFuncInfo_Method(b *testing.B) {
	v := testStruct{}
	for i := 0; i < b.N; i++ {
		GetFuncInfo(v.testMethod)
	}
}

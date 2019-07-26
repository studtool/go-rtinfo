package rtinfo

import (
	"testing"

	"github.com/franela/goblin"
)

func TestGetStructInfo(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("GetStructInfo()", func() {
		g.It("should return struct name for value", func() {
			v := testStruct{}
			info := GetStructInfo(v)

			g.Assert(info.NameWithPkg).Equal("rtinfo.testStruct")
		})
		g.It("should return struct name for pointer", func() {
			v := &testStruct{}
			info := GetStructInfo(v)

			g.Assert(info.NameWithPkg).Equal("rtinfo.testStruct")
		})
	})
}

func BenchmarkGetStructInfo_Value(b *testing.B) {
	v := testStruct{}
	for i := 0; i < b.N; i++ {
		GetStructInfo(v)
	}
}

func BenchmarkGetStructInfo_Pointer(b *testing.B) {
	v := &testStruct{}
	for i := 0; i < b.N; i++ {
		GetStructInfo(v)
	}
}

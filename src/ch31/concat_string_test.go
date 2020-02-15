package concat

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

const numbers = 100

// 字符串连接 Sprintf
func BenchmarkSprintf(b *testing.B)  {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var s string
		for j := 0; j < numbers; j++ {
			s = fmt.Sprintf("%v%v", s, j)
		}
	}
	b.StopTimer()
}

// 字符串连接 Builder
func BenchmarkStringBuilder(b *testing.B)  {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < numbers; j++ {
			builder.WriteString(strconv.Itoa(j))
		}
		_ = builder.String()
	}
	b.StopTimer()
}

// 字符串连接 bytes
func BenchmarkStringBuf(b *testing.B)  {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for j := 0; j < numbers; j++ {
			buf.WriteString(strconv.Itoa(j))
		}
		_ = buf.String()
	}
	b.StopTimer()
}
// 字符串连接 bytes
func BenchmarkStringAdd(b *testing.B)  {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var s string
		for j := 0; j < numbers; j++ {
			s += strconv.Itoa(j)
		}
	}
	b.StopTimer()
}
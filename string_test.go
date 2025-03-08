package lines

import (
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	ans := []string{
		"hoge",
		"fuga",
		"piyo",
		"",
		"foo bar baz",
	}
	s := strings.Join(ans, "\n")

	i := 0
	for line := range String(s) {
		if line != ans[i] {
			t.Fatalf("Expected %v, but got %v", ans[i], line)
		}
		i++
	}
	if i != 5 {
		t.Fatalf("Expected 5 lines, but there are only %d lines", i)
	}
}

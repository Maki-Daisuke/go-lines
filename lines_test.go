package lines

import (
	"os"
	"testing"
)

func TestLines(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	ans := []string{
		"hoge",
		"fuga",
		"piyo",
		"",
		"foo bar baz",
	}
	lines, e := Lines(file)
	i := 0
	for line := range lines {
		if line != ans[i] {
			t.Fatalf("Expected %v, but got %v", ans[i], line)
		}
		i++
	}
	if i != 5 {
		t.Fatal("Expected 5 lines, but there are only %d lines", i)
	}
	err = <-e
	if err != nil {
		t.Fatal(err)
	}
}

package lines

import (
	"os"
	"testing"
)

func TestLinesWithError(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	ans := []string{
		"hoge",
		"fuga",
		"piyo",
		"",
		"foo bar baz",
	}
	lines, e := LinesWithError(file)
	i := 0
	for line := range lines {
		if line != ans[i] {
			t.Fatalf("Expected %v, but got %v", ans[i], line)
		}
		i++
	}
	if i != 5 {
		t.Fatalf("Expected 5 lines, but there are only %d lines", i)
	}
	err = <-e
	if err != nil {
		t.Fatal(err)
	}
}

func TestLines(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	ans := []string{
		"hoge",
		"fuga",
		"piyo",
		"",
		"foo bar baz",
	}
	i := 0
	for line := range Lines(file) {
		if line != ans[i] {
			t.Fatalf("Expected %v, but got %v", ans[i], line)
		}
		i++
	}
	if i != 5 {
		t.Fatalf("Expected 5 lines, but there are only %d lines", i)
	}
}

package lines

import (
	"os"
	"testing"
)

func TestReader(t *testing.T) {
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
	var readErr error
	func() {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					readErr = err
				} else {
					panic(r)
				}
			}
		}()
		for line := range Reader(file) {
			if line != ans[i] {
				t.Fatalf("Expected %v, but got %v", ans[i], line)
			}
			i++
		}
	}()
	if readErr != nil {
		t.Fatal(readErr)
	}
	if i != 5 {
		t.Fatalf("Expected 5 lines, but there are only %d lines", i)
	}
}

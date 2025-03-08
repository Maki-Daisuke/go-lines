// Copyright 2015 Daisuke (yet another) Maki. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package lines makes it a bit easier to read lines from text files in Go.

# Description

There are so many ways to read lines in Go! But, I wanted to write less lines of
code as possible. I think life is too short to write lines of code to read lines.)

For example, you need to write code like the following to read line from STDIN:

	import (
	  "bufio"
	  "io"
	  "os"
	)

	func main(){
	  r := bufio.NewReader(os.Stdin)
	  line := ""
	  for {
	    l, isPrefix, err := r.ReadLine()
	    if err == io.EOF {
	      break
	    } else if err != nil {
	      panic(err)
	    }
	    line += l
	    if !isPrefix {
	      do_something_with(line)  // this is what really I want to do.
	      line = ""
	    }
	  }
	}

With this package and Go 1.23's range over func feature, you can write like this:

	import (
	  "os"
	  . "github.com/Maki-Daisuke/go-lines"
	)

	func main(){
	  for line := range lines.Reader(os.Stdin) {
	    do_something_with(line)
	  }
	}

Yay! It's much less lines of code and much cleaner!

If an error occurs during reading, the Reader function will panic.
If you want to handle errors, you can use a recover:

	import (
	  "fmt"
	  "os"
	  . "github.com/Maki-Daisuke/go-lines"
	)

	func main(){
	  defer func() {
	    if err := recover(); err != nil {
	        // Handle the error
	    }
	  }()

	  for line := range lines.Reader(os.Stdin) {
	    do_something_with(line)
	  }
	}

It's still less lines of code, isn't it?
*/
package lines

import (
	"bufio"
	"io"
)

// `Reader` converts a `io.Reader` to a func that can be used with range
// to iterate over lines. If an error occurs during reading, the function will panic.
// With Go 1.23's range over func feature, you can use it like:
//
//	for line := range Reader(reader) {
//	  do_something_with(line)
//	}
func Reader(r io.Reader) func(yield func(string) bool) {
	br := bufio.NewReader(r)
	return func(yield func(string) bool) {
		linebuf := ""
		for {
			line, isPrefix, err := br.ReadLine()
			if err == io.EOF {
				return
			} else if err != nil {
				panic(err)
			}
			linebuf += string(line)
			if !isPrefix {
				if !yield(linebuf) {
					return
				}
				linebuf = ""
			}
		}
	}
}

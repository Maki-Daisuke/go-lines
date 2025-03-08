# go-lines

    import "github.com/Maki-Daisuke/go-lines"

Package lines makes it a bit easier to read lines from text files in Go.


### Description

There are so many ways to read lines in Go! But, I wanted to write less lines of
code as possible. I think life is too short to write lines of code to read
lines.)

For example, you need to write code like the following to read line from STDIN:

```go
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
```

With go-lines package and Go 1.23's range over func feature, you can write like this:

```go
import (
  "os"
  "github.com/Maki-Daisuke/go-lines"
)

func main(){
  for line := range lines.Reader(os.Stdin) {
    do_something_with(line)
  }
}
```

Yay! It's much less lines of code and much cleaner!

If an error occurs during reading, the Reader function will panic.
If you want to handle errors, you can use a recover:

```go
import (
  "fmt"
  "os"
  "github.com/Maki-Daisuke/go-lines"
)

func main(){
  defer func() {
    if err := recover(); err != nil {
      // Handle the error
      fmt.Println("Error:", err)
    }
  }()
  
  for line := range lines.Reader(os.Stdin) {
    do_something_with(line)
  }
}
```

It's still less lines of code, isn't it?


## Usage

#### func Reader

```go
func Reader(r io.Reader) func(yield func(string) bool)
```
`Reader` converts a `io.Reader` to a func that can be used with range to iterate over lines. 
If an error occurs during reading, the function will panic. With Go 1.23's range over func 
feature, you can use it like:

```go
for line := range Reader(reader) {
  do_something_with(line)
}
```


## License

The Simplified BSD License (2-clause).
See [LICENSE](LICENSE) file also.


## Author

Daisuke (yet another) Maki

# go-lines

    import lines "github.com/Maki-Daisuke/go-lines"

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

With go-lines package, you can write like this:

```go
import (
  "os"
  . "github.com/Maki-Daisuke/go-lines"
)

func main(){
  for line := range Lines(os.Stdin) {
    do_something_with(line)
  }
}
```

Yay! It's much less lines of code!

Huh? How about error handling? Ok, you are not so lazy. You can use another
function `LinesWithError`:

```go
import (
  "os"
  . "github.com/Maki-Daisuke/go-lines"
)

func main(){
  lines, errs := LinesWithError(os.Stdin)
  for line := range lines {
    do_something_with(line)
  }
  err := <-errs
  if err != nil {
    panic(err)
  }
}
```

It's still less lines of code, isn't it?


## Usage

#### func  Lines

```go
func Lines(r io.Reader) <-chan string
```
`Lines` converts a `io.Reader` to a channel that generates a line for each
receipt. This is actually a shorthand of LinesWithError, ignoring errors.

#### func  LinesWithError

```go
func LinesWithError(r io.Reader) (lines <-chan string, err <-chan error)
```
`LinesWithError` converts a `io.Reader` to a channel that generates a line for
each receipt. If error occurs druing reading lines, it sends `error` to channel
`errs`. `errs` is sent at most one value. If there is no error, `err` receives
`nil`.


## License

The Simplified BSD License (2-clause).
See [LICENSE](LICENSE) file also.


## Author

Daisuke (yet another) Maki

package lines

import (
	"bufio"
	"io"
)

func Lines(r io.Reader) <-chan string {
	lines, _ := LinesWithError(r)
	return lines
}

func LinesWithError(r io.Reader) (<-chan string, <-chan error) {
	br := bufio.NewReader(r)
	chan_line := make(chan string, 0)
	chan_error := make(chan error, 1)
	go func() {
		defer func() {
			close(chan_line)
			close(chan_error)
		}()
		linebuf := ""
		for {
			line, isPrefix, err := br.ReadLine()
			if err == io.EOF {
				chan_error <- nil
				return
			} else if err != nil {
				chan_error <- err
				return
			}
			linebuf += string(line)
			if !isPrefix {
				chan_line <- linebuf
				linebuf = ""
			}
		}
	}()
	return chan_line, chan_error
}

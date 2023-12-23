package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	c := exec.Command("cat")
	stdin, _ := c.StdinPipe()
	stdout, _ := c.StdoutPipe()
	stderr, _ := c.StderrPipe()

	// go io.Copy(stdin, os.Stdin)
	go addPrefix("", stdin, os.Stdin)
	// go io.Copy(os.Stdout, stdout)
	go addPrefix("out", os.Stdout, stdout)
	// go io.Copy(os.Stderr, stderr)
	go addPrefix("err", os.Stderr, stderr)

	c.Run()
}

func addPrefix(prefix string, w io.Writer, r io.Reader) {
	s := bufio.NewScanner(r)
	if prefix != "" {
		prefix = prefix + ": "
	}
	for s.Scan() {
		b := []byte(fmt.Sprintf("%v%s\n", prefix, s.Text()))
		io.Copy(w, bytes.NewBuffer(b))
	}
}

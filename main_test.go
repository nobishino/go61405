package main_test

import (
	"fmt"

	"github.com/nobishino/go61405/coro"
)

const input = `"hello world"`

func ExampleParser() {
	var p parser
	p.Init()
	for _, b := range []byte(input) {
		switch p.Write(b) {
		case BadInput:
			fmt.Println("bad input")
		case Success:
			fmt.Println("done")
			return
		}
	}
	fmt.Println("ran out of input")
	// Output:
	// done
}

func parseQuoted(read func() byte) bool {
	if read() != '"' {
		return false
	}
	var c byte
	for c != '"' {
		c = read()
		if c == '\\' {
			read()
		}
	}
	return true
}

type Status int

const (
	NeedMoreInput Status = iota
	BadInput
	Success
)

type parser struct {
	resume func(byte) Status
}

func (p *parser) Init() {
	coparse := func(_ byte, yield func(Status) byte) Status {
		read := func() byte { return yield(NeedMoreInput) }
		if !parseQuoted(read) {
			return BadInput
		}
		return Success
	}
	p.resume = coro.New(coparse)
	p.resume(0)
}

func (p *parser) Write(c byte) Status {
	return p.resume(c)
}
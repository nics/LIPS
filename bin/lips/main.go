package main

import (
	"bufio"
	"flag"
	"fmt"
	. "github.com/nics/lips"
	"os"
)

func readFile(self *Interpreter, path string) {
	if _, e := self.ReadFile(path); e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
}

func repl(self *Interpreter) {
	in := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to LIPS, a misspelled LISP.")
	var expr, cell Cell
	var e error
	for {
		if expr, e = self.ReadExpression(in); e == nil {
			if cell, e = self.Eval(expr, self.Globals); e == nil {
				fmt.Println("=> " + Sexp(cell))
				continue
			}
		}
		fmt.Fprintln(os.Stderr, e)
		if _, is := e.(Error); !is {
			os.Exit(1)
		}
	}
}

func main() {
	self := NewInterpreter()

	flag.Parse()
	if len(flag.Args()) == 0 {
		repl(self)
	} else {
		for _, path := range flag.Args() {
			readFile(self, path)
		}
	}
}

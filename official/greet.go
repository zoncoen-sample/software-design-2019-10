package main

import (
	"errors"
	"fmt"
	"os"
	"plugin"
)

func main() {
	if len(os.Args) != 2 {
		exit(errors.New("invalid argument"))
	}
	if err := greet(os.Args[1]); err != nil {
		exit(err)
	}
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	os.Exit(1)
}

func greet(lang string) error {
	// Open plugin.
	p, err := plugin.Open(lang + ".so")
	if err != nil {
		return err
	}

	// Lookup a symbol named "Greet".
	v, err := p.Lookup("Greet")
	if err != nil {
		return err
	}

	f, ok := v.(func())
	if !ok {
		return errors.New(`Greet must be a "func()"`)
	}
	f()
	return nil
}

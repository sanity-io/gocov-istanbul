package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/axw/gocov"

	gocovistanbul "github.com/sanity-io/gocov-istanbul"
)

type Cover struct {
	Packages []*gocov.Package
}

func run() error {
	var cover Cover
	dec := json.NewDecoder(os.Stdin)
	if err := dec.Decode(&cover); err != nil {
		return fmt.Errorf("parsing stdin: %w", err)
	}

	c := gocovistanbul.NewConverter()
	for _, pkg := range cover.Packages {
		if err := c.AddPackage(pkg); err != nil {
			return fmt.Errorf("handling coverage from %s: %w", pkg.Name, err)
		}
	}

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(c.Finalize()); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error: %s\n", err)
	}
}

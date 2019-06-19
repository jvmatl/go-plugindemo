// +build linux darwin
package main

// This package is a plugin. Build it with `go build -buildmode=plugin -o shout.so`

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jvmatl/go-plugindemo/processors"
)

// ShoutProcessor will capitalize any byte slices passed in
type ShoutProcessor struct {
	configured    bool
	logEverything bool
}

// NewProcessor is more strongly typed, and a better way to go if you expect to have many plugins
func NewProcessor() processors.Processor {
	return &ShoutProcessor{}
}

// GenericNew is the quick and dirty way to do this, without needing the separate processors package
func GenericNew() interface{} {
	return &ShoutProcessor{}
}

// Init accepts configuration information for your processor object
func (p *ShoutProcessor) Init(config map[string]interface{}) error {
	var ok bool
	if p.logEverything, ok = config["log_everything"].(bool); !ok {
		return errors.New("invalid config")
	}

	p.configured = true
	return nil
}

// Process will take in a []byte and do something cool with it. :)
func (p *ShoutProcessor) Process(buf []byte) []byte {
	if p == nil || !p.configured {
		panic(fmt.Sprintf("Unconfigured %T!", p))
	}

	if p.logEverything {
		fmt.Printf("  Shouter got data: %v\n", buf)
	}

	return []byte(strings.ToUpper(string(buf)))
}

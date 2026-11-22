package main

import (
	"flag"

	"github.com/Kabisdox/task-3/internal/config"
	"github.com/Kabisdox/task-3/internal/parser"
	"github.com/Kabisdox/task-3/internal/sorter"
	"github.com/Kabisdox/task-3/internal/writer"
)

func main() {
	cfgPath := flag.String("config", "config.yaml", "path to configuration file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	if err != nil {
		panic(err)
	}

	// Dependency injection: main решает, какие реализации использовать.
	var p parser.Parser = parser.XMLParser{}
	var w writer.Saver = writer.JSONWriter{}

	items, err := p.Parse(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	sorter.ByValueDesc(items)

	if err := w.Save(cfg.OutputFile, items); err != nil {
		panic(err)
	}
}

package main

import (
	"flag"

	"github.com/jaimecabrito01/separador-arquivos-service/internal/cli"
	"github.com/jaimecabrito01/separador-arquivos-service/internal/watcher"
)

func main() {
	setup := flag.Bool("setup", false, "configure paths")
	flag.Parse()

	if *setup {
		cli.Init()
		return
	}

	watcher.Watcher()
}

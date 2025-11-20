package main

import (
	"flag"

	"github.com/jaimecabrito01/separador-arquivos-service/internal/tela"
	"github.com/jaimecabrito01/separador-arquivos-service/internal/watcher"
)

func main() {
	setup := flag.Bool("setup", false, "configure paths")
	flag.Parse()

	if *setup {
		tela.Init()
		return
	}

	watcher.Watcher()
}

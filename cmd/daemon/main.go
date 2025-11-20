package main

import (
	"github.com/jaimecabrito01/separador-arquivos-service/internal/tela"
	"github.com/jaimecabrito01/separador-arquivos-service/internal/watcher"
)

func main() {
	tela.Init()
	watcher.Watcher()
}

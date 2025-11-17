package main

import (
	"github.com/jaimecabrito01/separador-arquivos-service/entities"
	"github.com/jaimecabrito01/separador-arquivos-service/internal/watcher"
)


func main() {

entities.NewConfig("/home/jaime/Downloads","/home/jaime/Imagens","/home/jaime/Vídeos","/home/jaime/Músicas","/home/jaime/Documentos")
watcher.Watcher()

}

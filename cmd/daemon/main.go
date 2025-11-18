package main

import (
	"os"
	"path/filepath"

	"github.com/jaimecabrito01/separador-arquivos-service/entities"
	"github.com/jaimecabrito01/separador-arquivos-service/internal/watcher"
)

func main() {
	home, _ := os.UserHomeDir()
	

	pat  := entities.Paths {
		Downloads: filepath.Join(home, "Downloads"),
		Images: filepath.Join(home, "Imagens"),
		Videos: filepath.Join(home, "Vídeos"),
		Musics: filepath.Join(home, "Músicas"),
		Documents: filepath.Join(home, "Docuementos"),

	}

	
	entities.NewConfig(&pat)
	watcher.Watcher()

}

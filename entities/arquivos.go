package entities

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Paths struct {
	Downloads string `json:"download"`
	Images    string `json:"images"`
	Videos    string `json:"videos"`
	Musics    string `json:"musics"`
	Documents string `json:"documents"`
}

func NewConfig(path *Paths) {
	/* path := Paths{
		Downloads:  Downloads,
		Images:     images,
		Videos:    videos,
		Musics:     musics,
		Documents: documentos,
	}
	*/

	pathJson, err := json.Marshal(*path)
	if err != nil {
		fmt.Println("Erro")
	}
	ex, _ := os.Executable()
	exPath := filepath.Dir(ex)
	config := filepath.Join(exPath, "config.json")

	er := os.WriteFile(config, pathJson, 0o644)
	if err != nil {
		log.Fatal(er)
	}
}

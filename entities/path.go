package entities

import (
	"encoding/json"
	"fmt"
	"os"
)

type Paths struct {
	Downloads  string `json:"download"`
	images     string `json:"images"`
	videos     string `json:"videos"`
	musics     string `json:"musics"`
	documentos string `json:"documents"`
}

func NewConfig(Downloads string, images string, videos string, musics string, documentos string) {
	path := Paths{
		Downloads:  Downloads,
		images:     images,
		musics:     musics,
		documentos: documentos,
	}
	pathJson, err := json.Marshal(path)
	if err != nil {
		fmt.Println("Erro")

	}
	os.WriteFile("config.json", pathJson, 0644)

}

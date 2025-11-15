package entities

import (
	"encoding/json"
	"fmt"
	"os"
)

type Paths struct {
	Downloads  string
	images     string
	videos     string
	musics     string
	documentos string
}

func NewPath(Downloads string, images string, videos string, musics string, documentos string) (*Paths, error) {
	path := Paths{
		Downloads:  Downloads,
		images:     images,
		musics:     musics,
		documentos: documentos,
	}
	pathJson, err := json.Marshal(path)
	if err != nil {
		return &path, err

	}
	os.WriteFile("config.json", pathJson, 0644)
	return &path, err

}

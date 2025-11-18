package organizer

import (
	"encoding/json"
	"fmt"
	"path/filepath"
"io"
	"os"
)

func LoadConfig() (string, string, string, string, string, error) {

	data, err := os.ReadFile("/home/jaime/go-projects/daemon/config.json")
	if err != nil {
		return "", "", "", "", "", err
	}

	var m map[string]interface{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return "", "", "", "", "", err
	}

	downloads, _ := m["download"].(string)
	musics, _ := m["musics"].(string)
	videos, _ := m["videos"].(string)
	documents, _ := m["documents"].(string)
	images, _ := m["images"].(string)

	return downloads, musics, videos, documents, images, nil
}

func Move(srcPath string, destDir string) {
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		os.MkdirAll(destDir, 0755)
	}

	fileName := filepath.Base(srcPath)
	destPath := filepath.Join(destDir, fileName)

	srcFile, err := os.Open(srcPath)
	if err != nil {
		fmt.Printf("Erro ao abrir origem %s: %v\n", srcPath, err)
		return
	}
	defer srcFile.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		fmt.Printf("Erro ao criar destino %s: %v\n", destPath, err)
		return
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		fmt.Printf("Erro ao copiar para %s: %v\n", destPath, err)
		os.Remove(destPath)
		return
	}

	srcFile.Close()
	destFile.Close()

	err = os.Remove(srcPath)
	if err != nil {
		fmt.Printf("Erro ao remover origem %s: %v\n", srcPath, err)
		return
	}

	fmt.Println("Movido para:", destPath)
}
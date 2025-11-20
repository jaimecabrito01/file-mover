package entities

import (
	"encoding/json"
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

func NewConfig(path *Paths) error {
	data, err := json.MarshalIndent(path, "", "  ")
	if err != nil {
		return err
	}
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	dir := filepath.Join(configDir, "filemover")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	configPath := filepath.Join(dir, "config.json")

	er := os.WriteFile(configPath, data, 0o644)
	if err != nil {
		log.Fatal(er)
	}
	return nil
}

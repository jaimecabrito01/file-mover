package organizer

import (
	"encoding/json"
	"fmt"

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

	downloads, _ := m["downloads"].(string)
	musics, _ := m["musics"].(string)
	videos, _ := m["videos"].(string)
	documents, _ := m["documents"].(string)
	images, _ := m["images"].(string)

	return downloads, musics, videos, documents, images, nil
}


func Move(path string,new string){

	err :=os.Rename(path,new)
	if err != nil {
		fmt.Println(err)
	}
}
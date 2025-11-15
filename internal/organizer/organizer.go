package organizer

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/jaimecabrito01/separador-arquivos-service/entities"
	"os"
)


data,_:= os.ReadFile("../../config.json")
var m map[string]interface{}
json.Unmarshal(data,&m)

downloads:= m["downloads"]
musics := m["musics"]
videos := m["videos"]
documents := m["documents"]
images := m["images"]
 
watcher,err := fsnotify.NewWatcher()
if err != nil {
	log.Fatal(err)

}
defer watcher.Close()

go func(){
	
}


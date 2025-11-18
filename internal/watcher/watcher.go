package watcher

import (
	"fmt"
	"log"
	"strings"
	

	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/jaimecabrito01/separador-arquivos-service/entities"
	"github.com/jaimecabrito01/separador-arquivos-service/internal/organizer"
)

func Watcher() {

	downloads, musics, videos, documents, images, err := organizer.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(downloads)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)

	}

	defer watcher.Close()
	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&(fsnotify.Create|fsnotify.Rename) != 0 {
					if strings.HasSuffix(event.Name, ".crdownload") ||
						strings.HasSuffix(event.Name, ".part") {
						continue
					}
					
					fmt.Println("event: ", event)
					ext := filepath.Ext(event.Name)
					for _, extVideo := range entities.VideoExtensions {
						if ext == string(extVideo) {
							organizer.Move(event.Name, videos)
						}
					}
					for _, extDocs := range entities.DocumentExtensions {
						if ext == string(extDocs) {
							organizer.Move(event.Name, documents)
						}
					}
					for _, extMusic := range entities.MusicExtensions {
						if ext == string(extMusic) {
							organizer.Move(event.Name, musics)
						}
					}
					for _, extImages := range entities.ImageExtensions {
						if ext == string(extImages) {
							
							organizer.Move(event.Name, images)
						}
					}

				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
	err = watcher.Add(downloads)
	if err != nil {
		log.Fatal(err)
	}

	<-done
}

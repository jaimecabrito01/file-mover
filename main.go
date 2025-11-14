package main

import "fmt"

var downloads string
var images string
var videos string
var musics string
var documentos string

func main() {

	fmt.Println("Onde voce quer guardar os tipos de arquivo " +
		".pdf,.doc,.docx.odt, etc...\n")
	fmt.Scanln(&documentos)

	fmt.Println("Onde voce quer colocar os arquivos de video\n")
	fmt.Scanln(&videos)

	fmt.Println("Onde colcoar os arquivos de musica\n")
	fmt.Scanln(&musics)

	fmt.Println("onde colocar as imagens\n")
	fmt.Scanln(&images)

	fmt.Println("Onde esta a pasta de downloads\n")
	fmt.Scanln(&downloads)

}

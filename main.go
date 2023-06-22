package main

import (
	"fmt"

	"github.com/DEHbNO4b/convert_en_to_vai/domain"
)

func main() {
	files, _ := domain.SearchNewEnFiles()
	for _, f := range files {
		enStrokes, err := domain.ReadEnFile(f)
		if err != nil {
			fmt.Println(err)
		}
		domain.CreateVaiDirs(enStrokes)
		//domain.CreateVaiFiles(enStrokes)
	}

}

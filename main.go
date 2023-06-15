package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/DEHbNO4b/convert_en_to_vai/data"
)

var filenameTemplate = "public/*.csv"
var fileDir = "public/"

func main() {
	files, err := searchNewFiles(fileDir)
	if err != nil {
		fmt.Println(err)
	}
	if len(files) == 0 {
		return
	}
	//ans := make([]data.Stroke, 0, 100)
	file, err := os.Open(files[0])
	if err != nil {
		return
	}

	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = ';'
	var en []data.Stroke

	for {

		record, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		//fmt.Println(record)
		stroke, err := data.MakeStroke(record)
		if err != nil {
			fmt.Println(err)
			break
		}

		en = append(en, stroke)
	}
	fmt.Println(en[0])
}

func searchNewFiles(p string) ([]string, error) {

	var files []string
	filepath.WalkDir(p, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			matched, err := filepath.Match(filepath.FromSlash(filenameTemplate), path)
			if err != nil {
				return err
			}
			if matched {
				//читаем файлы совпавшие с заданной строкой в требуемой директории
				files = append(files, path)

			}
		}
		return nil
	})

	return files, nil
}

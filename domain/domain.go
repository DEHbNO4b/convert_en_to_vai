package domain

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/DEHbNO4b/convert_en_to_vai/data"
)

var filenameTemplate = "public/*.csv"
var fileDir = "public/"

func SearchNewEnFiles() ([]string, error) {

	var files []string
	filepath.WalkDir(fileDir, func(path string, d fs.DirEntry, err error) error {
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
func CreateVaiDirs(ens []data.StrokeEn) {
	for _, en := range ens {

		year := strconv.Itoa(en.Time.Year())
		y := filepath.Join("public", year)
		os.Mkdir(y, 0755)
		month := strconv.Itoa(int(en.Time.Month()))
		m := filepath.Join(y, month)
		os.Mkdir(m, 0755)
		day := strconv.Itoa(en.Time.Day())
		d := filepath.Join(m, day)
		os.Mkdir(d, 0755)
		hour := strconv.Itoa(en.Time.Hour())
		h := filepath.Join(d, hour)
		os.Mkdir(h, 0755)

		min := strconv.Itoa(en.Time.Minute())

		f := filepath.Join(h, min+".lf")

		file, _ := os.OpenFile(f, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		// writer := csv.NewWriter(file)
		// writer.Comma = '\t'
		writer := bufio.NewWriter(file)
		vai := buildVaiFromEn(en)
		writer.WriteString(vai.String())
		//writer.Write(vai.Slice())
		writer.Flush()
		file.Close()
	}
}
func CreateVaiFiles(ens []data.StrokeEn) {
	// vais := make([]data.StrokeVai, 0, 5000)
	// for _, en := range ens {
	// 	vai := buildVaiFromEn(en)
	// 	vais = append(vais, vai)
	// }
	// file :=os.OpenFile()
	// fmt.Println(vais[1])

}
func buildVaiFromEn(en data.StrokeEn) data.StrokeVai {
	vai := data.StrokeVai{
		Version:        "0",
		Year:           strconv.Itoa(en.Time.Year()),
		Month:          strconv.Itoa(int(en.Time.Month())),
		Day:            strconv.Itoa(en.Time.Day()),
		Hour:           strconv.Itoa(en.Time.Hour()),
		Minutes:        strconv.Itoa(en.Time.Minute()),
		Seconds:        strconv.Itoa(en.Time.Second()),
		Nano:           en.Nano,
		Latitude:       en.Lat,
		Longitude:      en.Long,
		Signal:         en.Current,
		Flashes:        "0",
		Sensors:        en.Sensors,
		Freedom:        "0",
		Ell_angle:      "0.00",
		Ell_semimajore: "0.00",
		Ell_semiminore: "0.00",
		Chi_square:     "0.00",
		RiseTime:       "0.0",
		DownTime:       "0.0",
		MaxRateRise:    "0.0",
		Cloud:          en.Cloud,
		Ind_angle:      "0",
		Ind_sygnal:     "0",
		Ind_tyme:       "0",
	}
	return vai
}
func MakeStrokeEn(record []string) (data.StrokeEn, error) {
	if len(record) != 8 {
		return data.StrokeEn{}, fmt.Errorf("invalid person slice: %v", record)
	}
	//определение типа разряда
	var cloud string
	if len(record[0]) == 12 {
		cloud = "0"
	} else if len(record[0]) == 13 {
		cloud = "1"
	} else {
		return data.StrokeEn{}, fmt.Errorf("invalid person slice: %v", record)
	}
	//парсинг времени
	layout := "2006-01-02 15:04:05"
	time, _ := time.Parse(layout, record[1])

	nano := record[2]
	lat := record[3]
	long := record[4]
	c, _ := strconv.Atoi(record[5])
	c = c / 1000
	current := strconv.Itoa(c)
	altitude := record[6]
	sensors := record[7]
	return data.StrokeEn{Cloud: cloud, Time: time, Nano: nano,
		Lat: lat, Long: long, Current: current, Altitude: altitude, Sensors: sensors}, nil

	//fmt.Println(record)
	//return Stroke{}, nil

}

func ReadEnFile(path string) ([]data.StrokeEn, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := csv.NewReader(file) //создаем ридер csv
	r.Comma = ';'
	en := make([]data.StrokeEn, 0, 5000)

	for { //читаем по строкам

		record, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		//fmt.Println(record)
		stroke, err := MakeStrokeEn(record)
		if err != nil {
			fmt.Println(err)
			break
		}

		en = append(en, stroke)
	}
	return en, nil
}

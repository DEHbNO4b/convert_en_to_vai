package data

import (
	"fmt"
	"time"
)

type Stroke struct {
	Claud    int
	Time     time.Time
	Nano     string
	Lat      string
	Long     string
	Current  string
	Altitude string
	Sensors  string
}

func MakeStroke(record []string) (Stroke, error) {
	if len(record) != 8 {
		return Stroke{}, fmt.Errorf("invalid person slice: %v", record)
	}
	//определение типа разряда
	var claud int
	if len(record[0]) == 12 {
		claud = 0
	} else if len(record[0]) == 13 {
		claud = 1
	} else {
		return Stroke{}, fmt.Errorf("invalid person slice: %v", record)
	}
	//парсинг времени
	layout := "2006-01-02 15:04:05"
	time, _ := time.Parse(layout, record[1])

	nano := record[2]
	lat := record[3]
	long := record[4]
	current := record[5]
	altitude := record[6]
	sensors := record[7]
	return Stroke{claud, time, nano, lat, long, current, altitude, sensors}, nil

	//fmt.Println(record)
	//return Stroke{}, nil

}

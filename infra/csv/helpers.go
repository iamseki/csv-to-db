package csv

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func getCSVReader(filename string) (*csv.Reader, *os.File) {
	csvfile, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	return r, csvfile
}

func getCSVNextRecord(r *csv.Reader) []string {
	record, err := r.Read()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		log.Fatal(err)
	}
	return record
}

func parseDate(input string, standardSampleLayout string) time.Time {
	date, err := time.Parse(standardSampleLayout, input)
	if err != nil {
		log.Fatalln(err)
	}
	return date
}

func parseFloat32(input string) float32 {
	n, err := strconv.ParseFloat(input, 32)
	if err != nil {
		log.Fatalln(err)
	}
	return float32(n)
}

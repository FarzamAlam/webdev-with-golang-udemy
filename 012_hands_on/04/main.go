package main

import (
	"encoding/csv"
	"log"
	"os"
	"text/template"
)

type share struct {
	Date     string
	Open     string
	High     string
	Low      string
	Close    string
	Volume   string
	AdjClose string
}

func main() {
	shares := readCSV("table.csv")
	tpl := template.Must(template.ParseFiles("tpl.gohtml"))
	err := tpl.Execute(os.Stdout, shares)
	if err != nil {
		log.Fatalln(err)
	}
}

func readCSV(fileName string) []share {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Unable to read file : ", err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse file as CSV : ", err)
	}
	var shares []share
	for _, record := range records {
		var share share
		share.Date = record[0]
		share.Open = record[1]
		share.High = record[2]
		share.Low = record[3]
		share.Close = record[4]
		share.Volume = record[5]
		share.AdjClose = record[6]
		shares = append(shares, share)
	}
	return shares
}

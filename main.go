package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

type PinoutRecord struct {
	Pin    string `json:"pin"`
	Signal string `json:"signal"`
}

type PortRecord struct {
	Port    string         `json:"port"`
	Records []PinoutRecord `json:"signals"`
}

func createPortMap(data [][]string) map[string][]PinoutRecord {
	var portMap = make(map[string][]PinoutRecord)

	for i, line := range data {
		if i > 0 {
			var record PinoutRecord

			for j, field := range line {
				switch j {
				case 1:
					record.Pin = field
				case 3:
					record.Signal = field
				}
			}
			if record.Pin != "" && record.Signal != "" {
				port := "GPIO" + string(record.Pin[1])

				if _, ok := portMap[port]; !ok {
					newRecList := []PinoutRecord{record}

					portMap[port] = newRecList
				} else {
					portMap[port] = append(portMap[port], record)
				}
			}
		}
	}
	return portMap
}

func main() {
	var (
		csvFileName string
		outFileName string
	)

	flag.StringVar(&csvFileName, "f", "", "Specify a .csv file to read.")
	flag.StringVar(&outFileName, "o", "output.json", "Specify an output .json file")
	flag.Parse()

	if csvFileName == "" {
		log.Fatal("csvFileName is empty!")
	}
	f, err := os.Open(csvFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	portMap := createPortMap(data)

	jsonData, err := json.MarshalIndent(portMap, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(outFileName, jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("success! wrote file to %s\n", outFileName)
}

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "Adib Hauzan Sofyan\n" +
		"Wati Komalasari, Gilang, Rinat\n" +
		"Joko Widodo, Prabowo, Anie"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
}

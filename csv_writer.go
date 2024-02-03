package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writter := csv.NewWriter(os.Stdout)

	_ = writter.Write([]string{"Adib", "Hauzan", "Sofyan"})
	_ = writter.Write([]string{"Wato", "Komala", "Sari"})
	_ = writter.Write([]string{"Gilang", "Aditia", "Saputra"})

	writter.Flush()

}

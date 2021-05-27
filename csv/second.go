package second

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type cordinates struct {
	State     string `json:"state"`
	District  string `json:"district"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func main() {

	filepath := "./csv/LatLong.csv"
	openfile, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Error occured , cannot open:/n", err)
	}

	fileData, err := csv.NewReader(openfile).ReadAll()
	if err != nil {
		log.Fatal("Error occured , cannot Read:/n", err)
	}
	fmt.Println(fileData)
}

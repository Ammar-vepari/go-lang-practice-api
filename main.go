package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Rto struct {
	RegNo string `json:"regNo"`
	Place string `json:"place"`
	State string `json:"state"`
}

var vehicleNo []Rto
var VehicleMap map[string]map[string]string

var portNumber string = ":8000"

func getData(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vehicleNo)

}

func getDataByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println("\n", params)
	data := params["id"]
	fmt.Printf("value is:%v of type: %T\n", data, data)
	value := VehicleMap[data]

	json.NewEncoder(w).Encode(value)

}

func CreateData(w http.ResponseWriter, r *http.Request) {

}

func UpdateData(w http.ResponseWriter, r *http.Request) {

}

func DeleteData(w http.ResponseWriter, r *http.Request) {

}

func main() {

	//vechileNo := []Rto{}
	filepath := "./csv/rto.csv"
	VehicleMap = make(map[string]map[string]string)

	openfile, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Error occured , cannot open:/n", err)
	}

	fileData, err := csv.NewReader(openfile).ReadAll()
	if err != nil {
		log.Fatal("Error occured , cannot Read:/n", err)
	}

	for _, value := range fileData {
		//fmt.Printf("At index: %d the value is %s\n", i, value)
		p := Rto{value[0], value[1], value[2]}
		vehicleNo = append(vehicleNo, p)
		VehicleMap[value[0]] = make(map[string]string)
		VehicleMap[value[0]][value[1]] = value[2]
	}
	//fmt.Println(vechileNo)

	r := mux.NewRouter()

	r.HandleFunc("/api/NamePlate", getData).Methods("GET")
	r.HandleFunc("/api/NamePlate/{id}", getDataByID).Methods("GET")
	r.HandleFunc("/api/NamePlate", CreateData).Methods("Post")
	r.HandleFunc("/api/NamePlate/{id}", UpdateData).Methods("Put")
	r.HandleFunc("/api/NamePlate/{id}", DeleteData).Methods("GET")
	fmt.Printf("Running on portNumber %v\n", portNumber)
	log.Fatal(http.ListenAndServe(portNumber, r))
}

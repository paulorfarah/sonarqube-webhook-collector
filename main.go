package main

import (
	"os"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"net/http"
	"github.com/gorilla/mux"
)

//type Analysis struct {
//	ID	string	`json:"taskId,omitempty"`
//	ServerUrl string `json:"serverUrl,omitempty"`
//	Status	string `json:"status,omitempty"`
//	AnalysedAt string `json:"analysedAt,omitempty"`
//	Revision string	`json:"revision,omitempty"`
//	Project  Project

//}

//type Project struct {
//	Key string `json:"project"`
//	Name string `json:"name"`
//	Url  string `json:url"`
//}

func ParseJson(w http.ResponseWriter, r *http.Request) {
	var result map[string]interface{}
	_ = json.NewDecoder(r.Body).Decode(&result)
	file, _ := json.MarshalIndent(result, "", " ")
	fmt.Println(result)
	fileName := time.Now().Format("20060102150405") 
	jsonFile, err := os.Create("./" + fileName + ".json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	jsonFile.Write(file)
	jsonFile.Close()
}

func main() {
	fmt.Println("running on port 8001...")
	router := mux.NewRouter()
	//router.HandleFunc("/sonarqube/", CreateAnalysis).Methods("POST")
	router.HandleFunc("/sonarqube/", ParseJson).Methods("POST")
	log.Fatal(http.ListenAndServe(":8001", router))
}


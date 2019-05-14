package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"encoding/json"

	"./server"
	"./persistence"
	"github.com/gorilla/mux"
)
var (
	Err      = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
)
func main() {
	// Router
	router := mux.NewRouter()

	router.HandleFunc("/server/{domain}", GetServerByDomain).Methods("GET")

	fmt.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// GetServerByDomain endpoint to read data filter by domain
func GetServerByDomain(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	domain := vars["domain"]
	db := persistence.SetupDB()
	defer db.Close()

	response, err := server.GetDataAPIServer(db, domain)

	if err != nil {
		sendInternalServerError(err, w)
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		sendInternalServerError(err, w)
	}
	sendOkResponse(jsonData, w)

}

func sendBadServerError(err error, w http.ResponseWriter) {
	Err.Printf("ApiTruora - in %v", err)
	http.Error(w, err.Error(), http.StatusBadRequest)
}
func sendInternalServerError(err error, w http.ResponseWriter) {
	Err.Printf("ApiTruora - in %v", err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func sendOkResponse(jsonData []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
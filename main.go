package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"encoding/json"

	"./server"
	"./utils"
	"./persistence"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)
var (
	Err      = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
)
func main() {
    // Enviroment variables validation
	err := utils.ValidateEnvVars()
	if err != nil {
		log.Fatal(err)
	}
		
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Get("/server/{domain}", GetServerByDomain)
	r.Get("/servers", GetServers)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	fmt.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// GetServerByDomain endpoint to read data filter by domain
func GetServerByDomain(w http.ResponseWriter, request *http.Request) {
	domain := chi.URLParam(request, "domain")
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
// GetServers return all searchs
func GetServers(w http.ResponseWriter, request *http.Request) {
	db := persistence.SetupDB()
	defer db.Close()

	response, err := server.GetItems(db)

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
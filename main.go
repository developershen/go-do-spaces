package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/minio/minio-go"
	"io"
	"log"
	"net/http"
)

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/getBuckets", getBuckets).Methods("GET")
	router.HandleFunc("/getObject/{name}", getObject).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
func getBuckets(w http.ResponseWriter, r *http.Request) {
	accessKey :="QM7PBPMLNDBAVYO4B46K"
	secKey := "dKej1v9rNnQqZxN/x6G4BLEhYbtcuvki0L45UUhh+ok"
	endpoint := "nyc3.digitaloceanspaces.com"
	//spaceName := "workbench-space" // Space names must be globally unique
	ssl := true

	// Initiate a client using DigitalOcean Spaces.
	client, err := minio.New(endpoint, accessKey, secKey, ssl)
	// List all Spaces.
	spaces, err := client.ListBuckets()
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(spaces)

}

func getObject(w http.ResponseWriter, r *http.Request) {
	accessKey :="QM7PBPMLNDBAVYO4B46K"
	secKey := "dKej1v9rNnQqZxN/x6G4BLEhYbtcuvki0L45UUhh+ok"
	endpoint := "nyc3.digitaloceanspaces.com"
	spaceName := "workbench-space" // Space names must be globally unique
	ssl := true

	// Initiate a client using DigitalOcean Spaces.
	client, err := minio.New(endpoint, accessKey, secKey, ssl)
	params := mux.Vars(r)
	objectName  := params["name"]
	// List all Spaces.
	spaces, err := client.GetObject(spaceName, objectName, minio.GetObjectOptions{} )
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(w, spaces)


}

package main
import (
	"fmt"
	"log"
	"github.com/minio/minio-go"
)

func main() {
	accessKey :="QM7PBPMLNDBAVYO4B46K"
	secKey := "dKej1v9rNnQqZxN/x6G4BLEhYbtcuvki0L45UUhh+ok"
	endpoint := "nyc3.digitaloceanspaces.com"
	//spaceName := "workbench-space" // Space names must be globally unique
	ssl := true

	// Initiate a client using DigitalOcean Spaces.
	client, err := minio.New(endpoint, accessKey, secKey, ssl)
	if err != nil {
		log.Fatal(err)
	}
	// List all Spaces.
	spaces, err := client.ListBuckets()
	if err != nil {
		log.Fatal(err)
	}
	for _, space := range spaces {
		fmt.Println(space.Name)
	}

}
package main

import (
	"encoding/json"
	"github.com/thecalcaholic/mc-controlcenter/api"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/seeruk/minecraft-rcon/rcon"
)

type config struct {
	host     string
	port     int
	password string
}

func main() {
	config := loadConfig("./config.json")

	client, err := rcon.NewClient(config.host, config.port, config.password)
	if err != nil {
		log.Fatalf("Error initializing rcon client!\n%s", err)
	}

	http.HandleFunc("/", pageIndex)
	http.HandleFunc("/api/server/", api.Server)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func pageIndex(w http.ResponseWriter, r *http.Request) {

}

func loadConfig(path string) config {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	var cfg config

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("Error parsing config: %s", err)
	}

	return cfg
}

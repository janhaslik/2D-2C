package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var mysqlServiceAddr string = "http://localhost:8001"
var mongodbServiceAddr string = "http://localhost:8002"

var httpClientMySql httpClient = HttpClient(mysqlServiceAddr)
var httpClientMongoDb httpClient = HttpClient(mongodbServiceAddr)

type Api struct {
	port string
}

func Server(portNumber string) Api {
	return Api{
		port: portNumber,
	}
}

func (api *Api) Start() {

	http.HandleFunc("/benchmark", handleBenchmark)

	err := http.ListenAndServe(api.port, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func handleBenchmark(w http.ResponseWriter, r *http.Request) {
	benchmarks := make([]Benchmark, 0)

	seeder := NewSeeder()

	data, err := seeder.SeedShips()

	if err != nil {
		return
	}

	// Benchmark SaveShips
	handleBenchmarkSaveShips(data, &benchmarks)

	// Benchmark GetShips
	handleBenchmarkGetShips(&benchmarks)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(benchmarks)
}

func handleBenchmarkSaveShips(data []Ship, benchmarks *[]Benchmark) {
	// Benchmark SaveShips
	startSaveShipsMySqL := time.Now()
	httpClientMySql.SaveShips(data)
	endSaveShipsMySql := time.Now()

	saveShipsMySqlElapsed := endSaveShipsMySql.Sub(startSaveShipsMySqL)

	startSaveShipsMongoDb := time.Now()
	httpClientMongoDb.SaveShips(data)
	endSaveShipsMongoDb := time.Now()

	saveShipsMongoDbElapsed := endSaveShipsMongoDb.Sub(startSaveShipsMongoDb)

	*benchmarks = append(*benchmarks, Benchmark{
		Size:      1000,
		Objective: "SaveShips",
		Mysql:     saveShipsMySqlElapsed,
		Mongodb:   saveShipsMongoDbElapsed,
	})
}

func handleBenchmarkGetShips(benchmarks *[]Benchmark) {
	startGetShipsMySqL := time.Now()
	httpClientMySql.GetShips()
	endGetShipsMySql := time.Now()

	getShipsMySqlElapsed := endGetShipsMySql.Sub(startGetShipsMySqL)

	startGetShipsMongoDb := time.Now()
	httpClientMongoDb.GetShips()
	endGetShipsMongoDb := time.Now()

	getShipsMongoDbElapsed := endGetShipsMongoDb.Sub(startGetShipsMongoDb)

	*benchmarks = append(*benchmarks, Benchmark{
		Size:      1000,
		Objective: "GetShips",
		Mysql:     getShipsMySqlElapsed,
		Mongodb:   getShipsMongoDbElapsed,
	})
}

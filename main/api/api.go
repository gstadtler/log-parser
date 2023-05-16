package api

import (
	"fmt"
	"net/http"
	"main/reader"
	"encoding/json"
	"errors"
	"os"
)

func HandleRequests() {
	http.HandleFunc("/log", getLog)
	http.HandleFunc("/stats", getLogStats)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func getLogStats(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Println("got /stats request")

	var data = reader.GetStatsData()

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func getLog(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Println("got /log request")

	var data = reader.GetEventLog()

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
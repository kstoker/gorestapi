package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"runtime"

	logger "gorestapi/logger"
	config "gorestapi/config"
	"gorestapi/database"
	"gorestapi/auth"
	"gorestapi/mededelingen"
	"gorestapi/links"

	"github.com/gorilla/mux"
)

const (
	VERSION = "0.01"
)
	
func main() {
	// put startup message in the logfile
	logger.Log.Printf("Server v%s pid=%d started with processes: %d", VERSION, os.Getpid(),runtime.GOMAXPROCS(runtime.NumCPU()))

	// Init router
	r := mux.NewRouter().StrictSlash(true)

	// Add handlers to the routers
	r.HandleFunc("/api/shutdown", shutdown)
	auth.AddToRouter(r)
	mededelingen.AddToRouter(r)
	links.AddToRouter(r)

	port := strconv.Itoa(config.RouterPort)
	fmt.Println("Listening on port " +  port + "...")
	logger.Log.Fatal(http.ListenAndServe(":" + port, r))
}

// Shutdown recieved
func shutdown(w http.ResponseWriter, r *http.Request) {
	database.CloseDB()
	fmt.Println("Server stopped")
	os.Exit(0)
}

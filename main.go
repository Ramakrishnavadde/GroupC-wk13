package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// MySQL database configuration
const (
	dbUser     = "root"     // MySQL username (change if different)
	dbPassword = "password" // MySQL password (change to your actual password)
	dbName     = "toronto_time"
)

// TimeResponse struct for JSON response
type TimeResponse struct {
	CurrentTime string `json:"current_time"`
}

var db *sql.DB

// Connect to MySQL database
func connectToDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", dbUser, dbPassword, dbName)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging the database: ", err)
	}
}

// Log the current time to the database
func logTimeToDatabase(timestamp time.Time) {
	query := "INSERT INTO time_log (timestamp) VALUES (?)"
	_, err := db.Exec(query, timestamp)
	if err != nil {
		log.Println("Error inserting time into database: ", err)
	}
}

// CurrentTimeHandler returns the current time in Toronto
func currentTimeHandler(w http.ResponseWriter, r *http.Request) {
	// Get the current time in Toronto (Eastern Time)
	loc, err := time.LoadLocation("America/Toronto")
	if err != nil {
		http.Error(w, "Error loading timezone", http.StatusInternalServerError)
		log.Println("Error loading timezone:", err)
		return
	}
	currentTime := time.Now().In(loc)

	// Log the time into the database
	logTimeToDatabase(currentTime)

	// Prepare the response
	response := TimeResponse{
		CurrentTime: currentTime.Format(time.RFC3339),
	}

	// Set the content type and send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Connect to the database
	connectToDB()
	defer db.Close()

	// Set up routing using Gorilla Mux
	r := mux.NewRouter()
	r.HandleFunc("/current-time", currentTimeHandler).Methods("GET")

	// Start the server
	port := ":8080"
	fmt.Println("Server running on http://localhost" + port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

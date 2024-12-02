# GroupC-wk13

Building a Go API for Currennt Toronto Time with MySQL Database Logging.

# Overview

This project is a Go-based API that provides the current Toronto time in JSON format. It logs each API request to a MySQL database, ensuring accurate time zonne adjustments to Toronto's local time.

# Features

- Returns the current Toronto time via the /current-time endpoint.
- Logs each API request timestamp to a MySQLQ database.
- Implement error handling for database and time zone operations.
- Bonus: Logging, data retrival endpoint, and Dockerized deployment.

# Requirements

- Go: Installed and  configured.
- MySQL: Installed and running.
- Docker: For Containerized deployment.

# Setup Instrutions

1. Clone the Repository
   Run the following commands in the bash:
   git clone <https://github.com/Ramakrishnavadde/GroupC-wk13.git>
   cd GroupC-wk13

2. Set Up MySQL Database 
   - Install and start MySQL.
   - Create a new database:
     Run the following commands in the sql:
     CREATE DATABASE toronto_time;

   - Create the time_log table:
     Run the following command in the sql:
     USE toronto_time;
     CREATE TABLE time_log (
        id INT AUTO_INCREMENT PRIMARY KEY,
        timestamp DATETIME NOT NULL
     );

3. Configure Environment Variables

4. Run the Application
   - Install dependencies:
     Run the following command in the bash:
     go mod tidy

    - Run the server:
      Run the following command in the bash:
      go run main.go

5. Test the API
   - Access the /current-time endpoint using curl or a browser
     Run the following command in the bash:
     curl http://localhost:8080/current-time


# Endpoints

1. GET/ current-time
   - Description: Returns the currnet time in Toronto in JSON format
   - Response Example:
     json
     {
        "time: "2024-12-01T12:34:56-05:00"
     }

2. GET/time-logs (Bonus Challenge)
   - Description: Retrieves all logged timestamps from the database.
   - Response Example:
     json
     [
        {"id": 1, "timestamp": "2024-12-01T12:34:56-05:00}
        {"id": 2, "timestamp": "2024-12-01T12:36:02-05:00}
     ]

# Dockerized Deployment

1. Build and run the Containers
   - Build the Docker images:
     Run the following command:
     docker-comppose build
   - start the containers:
     Run the following command:
     docker-compose up

2. Access the API
   - The API will be accessible at http://localhost:8080/current-time.



package main

import (
    "log"
    "net/http"
    "os"

    "sast/internal/db"
    "sast/internal/handlers"
)

func main() {
    dsn := os.Getenv("MYSQL_DSN")
    if dsn == "" {
        dsn = "root:password@tcp(localhost:3306)/sast"
    }

    database, err := db.New(dsn)
    if err != nil {
        log.Fatalf("db connection failed: %v", err)
    }
    defer database.Close()

    http.HandleFunc("/tasks", handlers.TaskHandler(database))
    http.HandleFunc("/reports", handlers.ReportHandler(database))

    log.Println("server listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

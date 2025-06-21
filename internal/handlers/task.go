package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"

    "sast/internal/models"
)

// TaskHandler handles creation of new tasks.
func TaskHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodPost:
            var t models.Task
            if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
            }
            if _, err := db.Exec("INSERT INTO tasks(name, created_at) VALUES (?, NOW())", t.Name); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            w.WriteHeader(http.StatusCreated)
        case http.MethodGet:
            rows, err := db.Query("SELECT id, name, created_at FROM tasks ORDER BY id DESC")
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            defer rows.Close()

            var tasks []models.Task
            for rows.Next() {
                var t models.Task
                if err := rows.Scan(&t.ID, &t.Name, &t.CreatedAt); err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
                tasks = append(tasks, t)
            }
            json.NewEncoder(w).Encode(tasks)
        default:
            http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        }
    }
}

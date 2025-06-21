package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"

    "sast/internal/models"
)

// ReportHandler returns reports for a given task.
func ReportHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        taskID := r.URL.Query().Get("task_id")
        if taskID == "" {
            http.Error(w, "task_id is required", http.StatusBadRequest)
            return
        }
        rows, err := db.Query("SELECT id, task_id, content FROM reports WHERE task_id=?", taskID)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer rows.Close()

        var reports []models.Report
        for rows.Next() {
            var rp models.Report
            if err := rows.Scan(&rp.ID, &rp.TaskID, &rp.Content); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            reports = append(reports, rp)
        }
        json.NewEncoder(w).Encode(reports)
    }
}

package server

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"marginalia/internal/storage"
	"marginalia/internal/store"
)

// api holds the dependencies shared by the HTTP API handlers.
type api struct {
	db          *sql.DB
	storagePath string
}

func (a *api) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// setupStatusResponse reports whether first-run setup has been completed.
type setupStatusResponse struct {
	Completed bool `json:"completed"`
}

// handleSetupStatus reports whether onboarding has been completed. Setup is
// considered complete once at least one user account exists.
func (a *api) handleSetupStatus(w http.ResponseWriter, r *http.Request) {
	completed, err := store.HasUsers(r.Context(), a.db)
	if err != nil {
		log.Printf("http: setup status: %v", err)
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal error"})
		return
	}
	writeJSON(w, http.StatusOK, setupStatusResponse{Completed: completed})
}

// handleStorage reports the configured library storage directory and the disk
// space available on the volume that contains it.
func (a *api) handleStorage(w http.ResponseWriter, r *http.Request) {
	info, err := storage.Stat(a.storagePath)
	if err != nil {
		log.Printf("http: storage: %v", err)
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal error"})
		return
	}
	writeJSON(w, http.StatusOK, info)
}

// writeJSON writes v as a JSON response with the given status code.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("http: encode response: %v", err)
	}
}

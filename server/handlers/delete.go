package handlers

import (
	"net/http"

	"github.com/zamachnoi/viewthis/data"
)

func DeleteAllSubmissionsHandler(w http.ResponseWriter, r *http.Request) {
    if err := data.DeleteAllSubmissions(); err != nil {
        http.Error(w, "Error deleting all submissions", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func DeleteAllFeedbackHandler(w http.ResponseWriter, r *http.Request) {
    if err := data.DeleteAllFeedback(); err != nil {
        http.Error(w, "Error deleting all feedback", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func DeleteAllQueuesHandler(w http.ResponseWriter, r *http.Request) {
    if err := data.DeleteAllQueues(); err != nil {
        http.Error(w, "Error deleting all queues", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func DeleteAllUsersHandler(w http.ResponseWriter, r *http.Request) {
    if err := data.DeleteAllUsers(); err != nil {
        http.Error(w, "Error deleting all users", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

func DeleteAllDataHandler(w http.ResponseWriter, r *http.Request) {
    if err := data.DeleteAllData(); err != nil {
        http.Error(w, "Error deleting all data", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}
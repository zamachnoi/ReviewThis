package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/zamachnoi/viewthis/data"
	"github.com/zamachnoi/viewthis/models"
)

// GetSubmissionsHandler fetches all submissions for a specific queue from the database
func GetSubmissionsByQueueIDHandler(w http.ResponseWriter, r *http.Request) {
    queueID, _ := strconv.Atoi(chi.URLParam(r, "queueID"))
    submissions, err := data.GetSubmissionsByQueueID(uint(queueID))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(submissions)
}

// CreateSubmissionHandler creates a new submission
func CreateSubmissionHandler(w http.ResponseWriter, r *http.Request) {
    var submission models.Submission
    json.NewDecoder(r.Body).Decode(&submission)

    createdSubmission, err := data.CreateSubmission(submission)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(createdSubmission)
}

// DeleteSubmissionByIDHandler deletes a submission by its ID
func DeleteSubmissionByIDHandler(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(chi.URLParam(r, "id"))
    err := data.DeleteSubmissionByID(uint(id))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Write([]byte("Submission deleted"))
}

// Update submission by ID
func UpdateSubmissionHandler(w http.ResponseWriter, r *http.Request) {
	var submission models.Submission
	json.NewDecoder(r.Body).Decode(&submission)
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	updatedSubmission, err := data.UpdateSubmission(uint(id), submission)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedSubmission)
}
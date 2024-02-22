package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/zamachnoi/viewthis/data"
	"github.com/zamachnoi/viewthis/models"
	"github.com/zamachnoi/viewthis/util"
)

// parseClaimsQueueId checks if the user is the owner of the queue
func parseClaimsQueueId(queueID uint, jwt string, limit int, page int) ([]models.Submission, error) {
    var submissions []models.Submission
    content := false



    if jwt != "" {
        _, claims, err := util.ParseJWTClaims(jwt)
        log.Println("Claims: ", claims)
        if err != nil {
            log.Println("Error parsing claims: ", err)
            return nil, err
        }

        ownerDbID, err := data.GetOwnerDbIDByQueueID(uint(queueID))
        if err != nil {
            log.Println("Error getting owner db id: ", err)
            return nil, err
        }

        if claims.DBID == ownerDbID {
            content = true
        }
    }

    submissions, err := data.GetSubmissionsByQueueID(queueID, limit, page, content)
    if err != nil {
        log.Println("Error getting submissions by queue id: ", err)
        return nil, err
    }
    return submissions, nil
}

// fetch all submissions for a specific queue
func GetSubmissionsByQueueIDHandler(w http.ResponseWriter, r *http.Request) {
    queueID, _ := strconv.Atoi(chi.URLParam(r, "queueID"))

    limit, page := util.ParseLimitAndPage(r)

    var jwtValue string = ""
    cookie, err := r.Cookie("_viewthis_jwt")
    if err != nil {
        log.Println("No JWT cookie found")
    } else {
        jwtValue = cookie.Value
    }

    log.Println("Getting submissions for queue: ", queueID, " and jwt: ", jwtValue);
    submissions, err := parseClaimsQueueId(uint(queueID), jwtValue, limit, page)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(submissions)
}

// GetSubmissionByIDHandler fetches a submission by its ID
func GetSubmissionByIDHandler(w http.ResponseWriter, r *http.Request) {
    submissionID, _ := strconv.Atoi(chi.URLParam(r, "id"))
    
    jwtCookie, err := r.Cookie("_viewthis_jwt")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    _, claims, err := util.ParseJWTClaims(jwtCookie.Value) // TODO: maybe catch error because of invalid token in the between checking expiry and this one where it could :/
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    submission, err := data.GetSubmissionByIDWithUserIDCheck(uint(submissionID), claims.DBID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(submission)
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


package data

import (
	"errors"
	"log"

	"github.com/zamachnoi/viewthis/lib"
	"github.com/zamachnoi/viewthis/models"
	"gorm.io/gorm"
)

// TODO: FIX THE ERRRecordNotFound thing (should be errors.Is)
func GetSubmissionsByQueueID(queueID uint, limit int, page int, content bool) ([]models.Submission, error) {
    log.Println("Getting all submissions for queue: ", queueID)
    var submissions []models.Submission
    db := lib.GetDB().Where("queue_id = ?", queueID)

    if page != 0 && limit != 0 {
        if page < 1 {
            page = 1
        }
        offset := (page - 1) * limit
        db = db.Offset(offset).Limit(limit)
    }

    if err := db.Find(&submissions).Error; err != nil {
        if !errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, err
        }
    }

    // if the content is private, then remove the content from the submissions
    if !content {
        for i := range submissions {
            if submissions[i].Private{
                submissions[i].Content = ""
            }
        }
    }

    return submissions, nil
}

func GetSubmissionByIDWithUserIDCheck(submissionID uint, userID uint) (models.Submission, error) {
    var submission models.Submission
    err := lib.GetDB().Preload("Feedbacks").Where("id = ?", submissionID).First(&submission).Error
    if err != nil {
        return models.Submission{}, err
    }

    if submission.UserID != userID {
        submission.Content = "" // Remove the content
    }
    return submission, nil
}



func CreateSubmission(submission models.Submission) (*models.Submission, error) {
	if err := lib.GetDB().Create(&submission).Error; err != nil {
		return nil, err
	}
	return &submission, nil
}

func DeleteSubmissionByID(id uint) error {
	if err := lib.GetDB().Delete(&models.Submission{}, id).Error; err != nil {
		return err
	}
	return nil
}

func UpdateSubmission(id uint, submission models.Submission) (*models.Submission, error) {
	if err := lib.GetDB().Model(&models.Submission{}).Where("id = ?", id).Updates(submission).Error; err != nil {
		return nil, err
	}
	return &submission, nil
}

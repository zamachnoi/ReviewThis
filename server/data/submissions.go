package data

import (
	"github.com/zamachnoi/viewthis/lib"
	"github.com/zamachnoi/viewthis/models"
	"gorm.io/gorm"
)

// TODO: FIX THE ERRRecordNotFound thing (should be errors.Is)
func GetSubmissionsByQueueIDWithContent(queueID uint) ([]models.Submission, error) {
    var submissions []models.Submission
    if err := lib.GetDB().Preload("Feedbacks").Where("queue_id = ?", queueID).Find(&submissions).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
        return nil, err
    }
    return submissions, nil
}

func GetSubmissionsByQueueIDNoContent(queueID uint) ([]models.Submission, error) {
    var submissions []models.Submission
    if err := lib.GetDB().Model(&models.Submission{}).Preload("Feedbacks").Where("queue_id = ?", queueID).Find(&submissions).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, nil
        }
        return nil, err
    }
	for i := range submissions {
		submissions[i].Content = ""
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

package data

import (
	"github.com/zamachnoi/viewthis/db"
	"github.com/zamachnoi/viewthis/models"
)

func GetSubmissionsByQueueID(queueID uint) ([]models.Submission, error) {
    var submissions []models.Submission
    if err := db.GetDB().Preload("Feedbacks").Where("queue_id = ?", queueID).Find(&submissions).Error; err != nil {
        return nil, err
    }
    return submissions, nil
}

func CreateSubmission(submission models.Submission) (*models.Submission, error) {
	if err := db.GetDB().Create(&submission).Error; err != nil {
		return nil, err
	}
	return &submission, nil
}

func DeleteSubmissionByID(id uint) error {
	if err := db.GetDB().Delete(&models.Submission{}, id).Error; err != nil {
		return err
	}
	return nil
}

func UpdateSubmission(id uint, submission models.Submission) (*models.Submission, error) {
	if err := db.GetDB().Model(&models.Submission{}).Where("id = ?", id).Updates(submission).Error; err != nil {
		return nil, err
	}
	return &submission, nil
}

package data

import (
	"github.com/zamachnoi/viewthis/db"
	"github.com/zamachnoi/viewthis/models"
)

func DeleteAllSubmissions() error {
    if err := db.GetDB().Delete(&models.Submission{}).Error; err != nil {
        return err
    }
    return nil
}

func DeleteAllFeedback() error {
    if err := db.GetDB().Delete(&models.Feedback{}).Error; err != nil {
        return err
    }
    return nil
}

func DeleteAllQueues() error {
    if err := db.GetDB().Delete(&models.Queue{}).Error; err != nil {
        return err
    }
    return nil
}

func DeleteAllUsers() error {
    if err := db.GetDB().Delete(&models.User{}).Error; err != nil {
        return err
    }
    return nil
}

func DeleteAllData() error {
    if err := DeleteAllFeedback(); err != nil {
        return err
    }
    if err := DeleteAllSubmissions(); err != nil {
        return err
    }
    if err := DeleteAllQueues(); err != nil {
        return err
    }
    if err := DeleteAllUsers(); err != nil {
        return err
    }
    return nil
}
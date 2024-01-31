package data

import (
	"errors"

	"github.com/zamachnoi/viewthis/db"
	"github.com/zamachnoi/viewthis/models"
	"gorm.io/gorm"
)

func GetAllQueues() ([]models.Queue, error) {
    var queues []models.Queue
    err := db.GetDB().Find(&queues).Error
    if err != nil {
        return nil, err
    }
    if len(queues) == 0 {
        return nil, errors.New("no queues found")
    }
    return queues, nil
}

func CreateQueue(queue models.Queue) (*models.Queue, error) {
	if err := db.GetDB().Create(&queue).Error; err != nil {
		return nil, err
	}
	return &queue, nil
}

// GetQueue returns the queue with the given ID and all submissions in it
func GetQueueByID(id uint) (*models.Queue, error) {
    var queue models.Queue
    if err := db.GetDB().Preload("Submissions").Find(&queue, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, nil
        }
        return nil, err
    }
    return &queue, nil
}

func GetQueueByName(name string) (*models.Queue, error) {
    var queue models.Queue
    err := db.GetDB().Where("name LIKE ?", name).First(&queue).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, nil
        }
        return nil, err
    }
    return &queue, nil
}

func ClearQueueByID(id uint) error {
    var queue models.Queue
    if err := db.GetDB().First(&queue, id).Error; err != nil {
        return err
    }
    if err := db.GetDB().Model(&queue).Association("Submissions").Clear(); err != nil {
        return err
    }
    return nil
}

func UpdateQueue(queue models.Queue) (*models.Queue, error) {
    if err := db.GetDB().Save(&queue).Error; err != nil {
        return nil, err
    }
    return &queue, nil
}

func DeleteQueue(id uint) error {
	var queue models.Queue
    if err := db.GetDB().First(&queue, id).Error; err != nil {
        return err
    }
    if err := db.GetDB().Delete(&queue).Error; err != nil {
        return err
    }
    return nil

}

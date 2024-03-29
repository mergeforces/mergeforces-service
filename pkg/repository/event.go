package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/mergeforces/mergeforces-service/pkg/models"
)

func ListEvents(db *gorm.DB) (models.Events, error) {
	events := make([]*models.Event, 0)
	if err := db.Find(&events).Error; err != nil {
		return nil, err
	}

	if len(events) == 0 {
		return nil, nil
	}

	return events, nil
}

func CreateEvent(db *gorm.DB, event *models.Event) (*models.Event, error) {
	if err := db.Create(event).Error; err != nil {
		return nil, err
	}

	return event, nil
}

func ReadEvent(db *gorm.DB, id uint) (*models.Event, error) {
	event := &models.Event{}
	if err := db.Where("id = ?", id).First(&event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func UpdateEvent(db *gorm.DB, Event *models.Event) error {
	if err := db.First(&models.Event{}, Event.ID).Update(Event).Error; err != nil {
		return err
	}

	return nil
}

func DeleteEvent(db *gorm.DB, id uint) error {
	event := &models.Event{}
	if err := db.Where("id = ?", id).Delete(&event).Error; err != nil {
		return err
	}

	return nil
}
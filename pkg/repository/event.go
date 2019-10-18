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

package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Events []*Event

type Event struct {
	gorm.Model
	Name         string
	Description  string
	Location     string
	Availability int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type EventDtos []*EventDto

type EventDto struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Location     string `json:"location"`
	Availability int    `json:"availability"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

func (e Event) ToDto() *EventDto {
	return &EventDto{
		ID:           e.ID,
		Name:         e.Name,
		Description:  e.Description,
		Location:     e.Location,
		Availability: e.Availability,
		CreatedAt:    e.CreatedAt.Format("2006-01-02T15:04:05-07:00"),
		UpdatedAt:    e.UpdatedAt.Format("2006-01-02T15:04:05-07:00"),
	}
}

func (es Events) ToDto() EventDtos {
	dtos := make([]*EventDto, len(es))
	for i, b := range es {
		dtos[i] = b.ToDto()
	}
	return dtos
}

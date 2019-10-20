package models

import (
	"github.com/jinzhu/gorm"
)

type Events []*Event

type Event struct {
	gorm.Model
	Name         string
	Description  string
	Location     string
	Availability int
}

type EventDtos []*EventDto

type EventDto struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Location     string `json:"location"`
	Availability int    `json:"availability"`
}

type EventForm struct {
	Name			string  `json:"name"`
	Description		string  `json:"description"`
	Location		string  `json:"location"`
	Availability	int  	`json:"availability"`
}

func (f *EventForm) ToModel() (*Event, error) {
	return &Event{
		Name:			f.Name,
		Description:	f.Description,
		Location:		f.Location,
		Availability:	f.Availability,
	}, nil
}

func (e Event) ToDto() *EventDto {
	return &EventDto{
		ID:           e.ID,
		Name:         e.Name,
		Description:  e.Description,
		Location:     e.Location,
		Availability: e.Availability,
	}
}

func (es Events) ToDto() EventDtos {
	dtos := make([]*EventDto, len(es))
	for i, b := range es {
		dtos[i] = b.ToDto()
	}
	return dtos
}

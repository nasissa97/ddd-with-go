package fundamentals

import (
	"errors"
	"time"
)

type Campaign struct {
	Id      string
	Title   string
	Goal    string
	EndDate time.Time
}

type CampaingMetadata struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	EndDate  string `json:"endDate"`
}

type MarketingCampaignModel struct {
	Id       string           `json:"ID"`
	Metadata CampaingMetadata `json:"metadata"`
}

func (m *MarketingCampaignModel) ToCampaign() (*Campaign, error) {
	if m.Id == "" {
		return nil, errors.New("campaign ID cannot be empty")
	}
	formattedDate, err := time.Parse("2006-01-02", m.Metadata.EndDate)
	if err != nil {
		return nil, errors.New("endData was not in a parsable format")
	}

	return &Campaign{
		Id:      m.Id,
		Title:   m.Metadata.Name,
		Goal:    m.Metadata.Category,
		EndDate: formattedDate,
	}, nil
}

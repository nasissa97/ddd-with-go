package fundamentals_test

import (
	"testing"

	"github.com/nasissa97/ddd-with-go/fundamentals"
)

func Test_ToCampaign(t *testing.T) {
	m := fundamentals.MarketingCampaignModel{
		Id: "abcdef-12345-ghij-6789",
		Metadata: fundamentals.CampaingMetadata{
			Name:     "my campaign",
			Category: "awsome",
			EndDate:  "2024-12-07",
		},
	}

	_, err := m.ToCampaign()
	if err != nil {
		t.Fatalf("err was not nil: %v", err)
	}
}

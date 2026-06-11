package health

import (
	"context"
	"strings"

	"github.com/flaviolpgjr/aletheia/backend/internal/domain"
)

const currentHospitalBaseline = 0

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) FindEvidence(
	ctx context.Context,
	text string,
) ([]domain.Evidence, error) {
	if !isHealthPromise(text) {
		return []domain.Evidence{}, nil
	}

	return []domain.Evidence{
		{
			Source:      "Ministério da Saúde / DATASUS",
			Title:       "Hospitais e Leitos",
			Description: "Base pública utilizada como referência para acompanhar hospitais e leitos no Brasil, com dados relacionados ao CNES.",
			URL:         "https://dadosabertos.saude.gov.br/dataset/hospitais-e-leitos",

			Indicator: "hospital_facilities",
			Value:     currentHospitalBaseline,
			Unit:      "hospitais",
			Reference: "CNES/DATASUS - Hospitais e Leitos",
		},
	}, nil
}

func isHealthPromise(text string) bool {
	text = strings.ToLower(text)

	keywords := []string{
		"hospital",
		"hospitais",
		"saúde",
		"sus",
		"leito",
		"leitos",
		"upa",
		"ubs",
		"posto de saúde",
	}

	for _, keyword := range keywords {
		if strings.Contains(text, keyword) {
			return true
		}
	}

	return false
}
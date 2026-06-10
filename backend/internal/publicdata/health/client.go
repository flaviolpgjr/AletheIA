package health

import (
	"context"
	"strings"

	"github.com/flaviolpgjr/aletheia/backend/internal/domain"
)

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
			Source:      "Ministério da Saúde / CNES",
			Title:       "Cadastro Nacional de Estabelecimentos de Saúde",
			Description: "Base oficial utilizada para acompanhar hospitais, leitos e estabelecimentos de saúde no Brasil.",
			URL:         "https://apidadosabertos.saude.gov.br/",
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
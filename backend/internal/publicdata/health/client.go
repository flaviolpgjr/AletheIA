package health

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/flaviolpgjr/aletheia/backend/internal/domain"
)

const (
	defaultBaseURL = "https://apidadosabertos.saude.gov.br"
	cacheTTL       = 24 * time.Hour
)

type Client struct {
	httpClient *http.Client
	baseURL    string

	cacheMu       sync.Mutex
	baselineValue int
	baselineUntil time.Time
}

type hospitalsAndBedsResponse struct {
	Hospitals []hospitalAndBedsItem `json:"hospitais_leitos"`
}

type hospitalAndBedsItem struct {
	HospitalName string `json:"nome_do_hospital"`
	State        string `json:"unidade_da_federacao_onde_fica_o_hospital"`
	City         string `json:"nome_do_municipio_onde_fica_o_hospital"`
	UnitType    string `json:"descricao_do_tipo_da_unidade"`
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 20 * time.Second,
		},
		baseURL: defaultBaseURL,
	}
}

func (c *Client) FetchHospitalFacilitiesBaseline(
	ctx context.Context,
) (domain.PublicDataBaseline, error) {
	totalHospitals, err := c.fetchHospitalsCount(ctx)
	if err != nil {
		return domain.PublicDataBaseline{}, err
	}

	return domain.PublicDataBaseline{
		Indicator:   "hospital_facilities",
		Scope:       "BR",
		Value:       float64(totalHospitals),
		Unit:        "hospitais",
		Source:      "DATASUS",
		Reference:   "CNES/DATASUS - Hospitais e Leitos",
		CollectedAt: time.Now(),
	}, nil
}

func (c *Client) FindEvidence(
	ctx context.Context,
	text string,
) ([]domain.Evidence, error) {
	if !isHealthPromise(text) {
		return []domain.Evidence{}, nil
	}

	totalHospitals, err := c.fetchHospitalsCount(ctx)
	if err != nil {
		return nil, err
	}

	return []domain.Evidence{
		{
			Source:      "Ministério da Saúde / DATASUS",
			Title:       "Hospitais e Leitos",
			Description: "Base pública utilizada como referência para acompanhar hospitais e leitos no Brasil, com dados relacionados ao CNES.",
			URL:         "https://apidadosabertos.saude.gov.br/assistencia-a-saude/hospitais-e-leitos",

			Indicator: "hospital_facilities",
			Value:     float64(totalHospitals),
			Unit:      "hospitais",
			Reference: "CNES/DATASUS - Hospitais e Leitos",
		},
	}, nil
}

func (c *Client) fetchHospitalsCount(ctx context.Context) (int, error) {
	if value, ok := c.getCachedBaseline(); ok {
		return value, nil
	}

	total, err := c.fetchHospitalsCountFromAPI(ctx)
	if err != nil {
		return 0, err
	}

	c.setCachedBaseline(total)

	return total, nil
}

func (c *Client) fetchHospitalsCountFromAPI(ctx context.Context) (int, error) {
	exists, err := c.hospitalExistsAtOffset(ctx, 0)
	if err != nil {
		return 0, err
	}

	if !exists {
		return 0, nil
	}

	low := 0
	high := 1

	for {
		exists, err := c.hospitalExistsAtOffset(ctx, high)
		if err != nil {
			return 0, err
		}

		if !exists {
			break
		}

		low = high
		high *= 2
	}

	left := low + 1
	right := high

	for left < right {
		mid := left + (right-left)/2

		exists, err := c.hospitalExistsAtOffset(ctx, mid)
		if err != nil {
			return 0, err
		}

		if exists {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left, nil
}

func (c *Client) hospitalExistsAtOffset(
	ctx context.Context,
	offset int,
) (bool, error) {
	items, err := c.fetchHospitalsPage(ctx, 1, offset)
	if err != nil {
		return false, err
	}

	return len(items) > 0, nil
}

func (c *Client) fetchHospitalsPage(
	ctx context.Context,
	limit int,
	offset int,
) ([]hospitalAndBedsItem, error) {
	url := fmt.Sprintf(
		"%s/assistencia-a-saude/hospitais-e-leitos?limit=%d&offset=%d",
		c.baseURL,
		limit,
		offset,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("health public data request failed with status %d", resp.StatusCode)
	}

	var result hospitalsAndBedsResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Hospitals, nil
}

func (c *Client) getCachedBaseline() (int, bool) {
	c.cacheMu.Lock()
	defer c.cacheMu.Unlock()

	if c.baselineValue > 0 && time.Now().Before(c.baselineUntil) {
		return c.baselineValue, true
	}

	return 0, false
}

func (c *Client) setCachedBaseline(value int) {
	c.cacheMu.Lock()
	defer c.cacheMu.Unlock()

	c.baselineValue = value
	c.baselineUntil = time.Now().Add(cacheTTL)
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
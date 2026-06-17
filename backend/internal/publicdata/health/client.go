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
	defaultLimit   = 20
	maxPages       = 500

	activeStatus = 1
)

var hospitalUnitTypeCodes = []int{5, 7}

type Client struct {
	httpClient *http.Client
	baseURL    string

	baselineRepository BaselineRepository

	cacheMu       sync.Mutex
	baselineValue int
	baselineUntil time.Time
}

type cnesEstablishmentsResponse struct {
	Establishments []cnesEstablishmentItem `json:"estabelecimentos"`
}

type cnesEstablishmentItem struct {
	CNESCode      int  `json:"codigo_cnes"`
	UnitTypeCode  int  `json:"codigo_tipo_unidade"`
	StateCode     int  `json:"codigo_uf"`
	CityCode      int  `json:"codigo_municipio"`
	DisableReason *int `json:"codigo_motivo_desabilitacao_estabelecimento"`
}

type BaselineRepository interface {
	FindByIndicatorAndScope(
		ctx context.Context,
		indicator string,
		scope string,
	) (*domain.PublicDataBaseline, error)
}

func NewClient(
	baselineRepository BaselineRepository,
) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 20 * time.Second,
		},
		baseURL:            defaultBaseURL,
		baselineRepository: baselineRepository,
	}
}

func (c *Client) FetchHospitalFacilitiesBaseline(
	ctx context.Context,
) (domain.PublicDataBaseline, error) {
	totalHospitals, err := c.fetchHospitalFacilitiesCount(ctx)
	if err != nil {
		return domain.PublicDataBaseline{}, err
	}

	return domain.PublicDataBaseline{
		Indicator:   "hospital_facilities",
		Scope:       "BR",
		Value:       float64(totalHospitals),
		Unit:        "hospitais",
		Source:      "DATASUS/CNES",
		Reference:   "CNES - Estabelecimentos de Saúde: codigo_tipo_unidade IN (5, 7), status=1",
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

	baseline, err := c.baselineRepository.FindByIndicatorAndScope(
		ctx,
		"hospital_facilities",
		"BR",
	)
	if err != nil {
		return nil, err
	}

	return []domain.Evidence{
		{
			Source:      "Ministério da Saúde / DATASUS",
			Title:       "CNES - Estabelecimentos de Saúde",
			Description: "Cadastro Nacional de Estabelecimentos de Saúde usado como referência para contar hospitais gerais e especializados ativos no Brasil.",
			URL:         "https://apidadosabertos.saude.gov.br/cnes/estabelecimentos",

			Indicator: baseline.Indicator,
			Value:     baseline.Value,
			Unit:      baseline.Unit,
			Reference: baseline.Reference,
		},
	}, nil
}

func (c *Client) fetchHospitalFacilitiesCount(ctx context.Context) (int, error) {
	if value, ok := c.getCachedBaseline(); ok {
		return value, nil
	}

	total, err := c.fetchHospitalFacilitiesCountFromAPI(ctx)
	if err != nil {
		return 0, err
	}

	c.setCachedBaseline(total)

	return total, nil
}

func (c *Client) fetchHospitalFacilitiesCountFromAPI(ctx context.Context) (int, error) {
	uniqueHospitals := make(map[int]struct{})

	for _, unitTypeCode := range hospitalUnitTypeCodes {
		for page := 0; page < maxPages; page++ {
			offset := page * defaultLimit

			items, err := c.fetchCNESEstablishmentsPage(
				ctx,
				unitTypeCode,
				defaultLimit,
				offset,
			)
			if err != nil {
				return 0, err
			}

			for _, item := range items {
				if !isValidHospitalFacility(item, unitTypeCode) {
					continue
				}

				uniqueHospitals[item.CNESCode] = struct{}{}
			}

			if len(items) == 0 {
				break
			}
		}
	}

	return len(uniqueHospitals), nil
}

func (c *Client) fetchCNESEstablishmentsPage(
	ctx context.Context,
	unitTypeCode int,
	limit int,
	offset int,
) ([]cnesEstablishmentItem, error) {
	url := fmt.Sprintf(
		"%s/cnes/estabelecimentos?codigo_tipo_unidade=%d&status=%d&limit=%d&offset=%d",
		c.baseURL,
		unitTypeCode,
		activeStatus,
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
		return nil, fmt.Errorf(
			"health public data request failed with status %d",
			resp.StatusCode,
		)
	}

	var result cnesEstablishmentsResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Establishments, nil
}

func isValidHospitalFacility(
	item cnesEstablishmentItem,
	expectedUnitTypeCode int,
) bool {
	if item.CNESCode == 0 {
		return false
	}

	if item.UnitTypeCode != expectedUnitTypeCode {
		return false
	}

	if item.DisableReason != nil {
		return false
	}

	return true
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
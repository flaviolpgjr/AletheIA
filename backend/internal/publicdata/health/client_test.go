package health

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/flaviolpgjr/aletheia/backend/internal/domain"
)

type fakeBaselineRepository struct {
	baseline *domain.PublicDataBaseline
	err      error
}

func (f *fakeBaselineRepository) FindByIndicatorAndScope(
	ctx context.Context,
	indicator string,
	scope string,
) (*domain.PublicDataBaseline, error) {
	if f.err != nil {
		return nil, f.err
	}

	return f.baseline, nil
}

func TestIsValidHospitalFacility(t *testing.T) {
	disableReason := 1

	tests := []struct {
		name                 string
		item                 cnesEstablishmentItem
		expectedUnitTypeCode int
		expected             bool
	}{
		{
			name: "valid hospital",
			item: cnesEstablishmentItem{
				CNESCode:     123,
				UnitTypeCode: 5,
			},
			expectedUnitTypeCode: 5,
			expected:             true,
		},
		{
			name: "invalid when cnes code is zero",
			item: cnesEstablishmentItem{
				CNESCode:     0,
				UnitTypeCode: 5,
			},
			expectedUnitTypeCode: 5,
			expected:             false,
		},
		{
			name: "invalid when unit type is different",
			item: cnesEstablishmentItem{
				CNESCode:     123,
				UnitTypeCode: 7,
			},
			expectedUnitTypeCode: 5,
			expected:             false,
		},
		{
			name: "invalid when hospital is disabled",
			item: cnesEstablishmentItem{
				CNESCode:      123,
				UnitTypeCode:  5,
				DisableReason: &disableReason,
			},
			expectedUnitTypeCode: 5,
			expected:             false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidHospitalFacility(
				tt.item,
				tt.expectedUnitTypeCode,
			)

			if result != tt.expected {
				t.Fatalf(
					"expected %v, got %v",
					tt.expected,
					result,
				)
			}
		})
	}
}

func TestFetchHospitalFacilitiesCountFromAPI(t *testing.T) {
	disableReason := 1

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			unitType := r.URL.Query().Get("codigo_tipo_unidade")
			offset := r.URL.Query().Get("offset")

			w.Header().Set("Content-Type", "application/json")

			if offset != "0" {
				_, _ = w.Write([]byte(`{
					"estabelecimentos": []
				}`))
				return
			}

			switch unitType {
			case "5":
				_, _ = w.Write([]byte(`{
					"estabelecimentos": [
						{
							"codigo_cnes": 100,
							"codigo_tipo_unidade": 5,
							"codigo_uf": 31,
							"codigo_municipio": 3106200,
							"codigo_motivo_desabilitacao_estabelecimento": null
						},
						{
							"codigo_cnes": 100,
							"codigo_tipo_unidade": 5,
							"codigo_uf": 31,
							"codigo_municipio": 3106200,
							"codigo_motivo_desabilitacao_estabelecimento": null
						},
						{
							"codigo_cnes": 0,
							"codigo_tipo_unidade": 5,
							"codigo_uf": 31,
							"codigo_municipio": 3106200,
							"codigo_motivo_desabilitacao_estabelecimento": null
						},
						{
							"codigo_cnes": 200,
							"codigo_tipo_unidade": 5,
							"codigo_uf": 31,
							"codigo_municipio": 3106200,
							"codigo_motivo_desabilitacao_estabelecimento": 1
						}
					]
				}`))

			case "7":
				_ = disableReason

				_, _ = w.Write([]byte(`{
					"estabelecimentos": [
						{
							"codigo_cnes": 300,
							"codigo_tipo_unidade": 7,
							"codigo_uf": 35,
							"codigo_municipio": 3550308,
							"codigo_motivo_desabilitacao_estabelecimento": null
						},
						{
							"codigo_cnes": 400,
							"codigo_tipo_unidade": 5,
							"codigo_uf": 35,
							"codigo_municipio": 3550308,
							"codigo_motivo_desabilitacao_estabelecimento": null
						}
					]
				}`))

			default:
				_, _ = w.Write([]byte(`{
					"estabelecimentos": []
				}`))
			}
		}),
	)
	defer server.Close()

	client := NewClient(nil)
	client.baseURL = server.URL

	total, err := client.fetchHospitalFacilitiesCountFromAPI(
		context.Background(),
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if total != 2 {
		t.Fatalf("expected 2 hospitals, got %d", total)
	}
}

func TestFetchHospitalFacilitiesBaseline(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			offset := r.URL.Query().Get("offset")

			w.Header().Set("Content-Type", "application/json")

			if offset != "0" {
				_, _ = w.Write([]byte(`{
					"estabelecimentos": []
				}`))
				return
			}

			_, _ = w.Write([]byte(`{
				"estabelecimentos": [
					{
						"codigo_cnes": 100,
						"codigo_tipo_unidade": 5,
						"codigo_uf": 31,
						"codigo_municipio": 3106200,
						"codigo_motivo_desabilitacao_estabelecimento": null
					}
				]
			}`))
		}),
	)
	defer server.Close()

	client := NewClient(nil)
	client.baseURL = server.URL

	baseline, err := client.FetchHospitalFacilitiesBaseline(
		context.Background(),
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if baseline.Indicator != "hospital_facilities" {
		t.Errorf("expected hospital_facilities, got %s", baseline.Indicator)
	}

	if baseline.Scope != "BR" {
		t.Errorf("expected BR, got %s", baseline.Scope)
	}

	if baseline.Value != 1 {
		t.Errorf("expected value 1, got %.2f", baseline.Value)
	}

	if baseline.Unit != "hospitais" {
		t.Errorf("expected hospitais, got %s", baseline.Unit)
	}

	if baseline.Source != "DATASUS/CNES" {
		t.Errorf("expected DATASUS/CNES, got %s", baseline.Source)
	}

	if baseline.CollectedAt.IsZero() {
		t.Fatal("expected collected_at to be set")
	}
}

func TestFindEvidenceReturnsEmptyWhenTextIsNotHealthPromise(t *testing.T) {
	client := NewClient(nil)

	evidence, err := client.FindEvidence(
		context.Background(),
		"Vou reduzir impostos federais em 20%",
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(evidence) != 0 {
		t.Fatalf("expected empty evidence, got %d", len(evidence))
	}
}

func TestFindEvidenceReturnsHealthEvidence(t *testing.T) {
	client := NewClient(&fakeBaselineRepository{
		baseline: &domain.PublicDataBaseline{
			Indicator:   "hospital_facilities",
			Scope:       "BR",
			Value:       5115,
			Unit:        "hospitais",
			Source:      "DATASUS/CNES",
			Reference:   "Cadastro Nacional de Estabelecimentos de Saúde (CNES) - Quantidade de hospitais ativos no Brasil",
			CollectedAt: time.Now(),
		},
	})

	evidence, err := client.FindEvidence(
		context.Background(),
		"Vou construir 100 hospitais públicos em 2 anos",
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(evidence) != 1 {
		t.Fatalf("expected 1 evidence, got %d", len(evidence))
	}

	if evidence[0].Indicator != "hospital_facilities" {
		t.Errorf("expected hospital_facilities, got %s", evidence[0].Indicator)
	}

	if evidence[0].Value != 5115 {
		t.Errorf("expected value 5115, got %.2f", evidence[0].Value)
	}

	if evidence[0].Unit != "hospitais" {
		t.Errorf("expected hospitais, got %s", evidence[0].Unit)
	}

	if evidence[0].Title != "Hospitais ativos registrados no CNES" {
		t.Errorf("unexpected title: %s", evidence[0].Title)
	}
}
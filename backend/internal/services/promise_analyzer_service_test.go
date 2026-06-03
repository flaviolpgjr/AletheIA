package services

import "testing"

func TestAnalyzeWhenPromiseMentionsTax(t *testing.T) {
	service := NewPromiseAnalyzerService()

	analysis := service.Analyze("reduzir imposto sobre combustível")

	if analysis.Score != 35 {
		t.Errorf("expected score 35, got %d", analysis.Score)
	}

	if len(analysis.Criteria) != 6 {
		t.Errorf("expected 6 criteria, got %d", len(analysis.Criteria))
	}

	if len(analysis.Risks) == 0 {
		t.Errorf("expected risks, got none")
	}
}

func TestAnalyzeWhenPromiseHasNoKnownKeywords(t *testing.T) {
	service := NewPromiseAnalyzerService()

	analysis := service.Analyze("melhorar a qualidade dos serviços públicos")

	if analysis.Score != 43 {
		t.Errorf("expected score 43, got %d", analysis.Score)
	}

	if len(analysis.Criteria) != 6 {
		t.Errorf("expected 6 criteria, got %d", len(analysis.Criteria))
	}

	if analysis.Risks[0] != "Dados públicos adicionais são necessários para aprofundar a avaliação." {
		t.Errorf("expected default risk, got %s", analysis.Risks[0])
	}
}
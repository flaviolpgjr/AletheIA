package services

import "testing"

func TestAnalyzeWhenPromiseMentionsTax(t *testing.T) {
	service := NewPromiseAnalyzerService()

	analysis := service.Analyze("reduzir imposto sobre combustível")

	if analysis.Score != 60 {
		t.Errorf("expected score 60, got %d", analysis.Score)
	}

	if len(analysis.Risks) == 0 {
		t.Errorf("expected risks, got none")
	}
}

func TestAnalyzeWhenPromiseHasNoKnownKeywords(t *testing.T) {
	service := NewPromiseAnalyzerService()

	analysis := service.Analyze("melhorar a qualidade dos serviços públicos")

	if analysis.Score != 75 {
		t.Errorf("expected score 75, got %d", analysis.Score)
	}

	if analysis.Risks[0] != "Dados adicionais necessários para avaliação" {
		t.Errorf("expected default risk, got %s", analysis.Risks[0])
	}
}

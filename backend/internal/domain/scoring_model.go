package domain

var ScoringModelV1 = []Criterion{
	{
		Key:    "clarity",
		Name:   "Clareza da promessa",
		Weight: 10,
	},
	{
		Key:    "measurability",
		Name:   "Mensurabilidade",
		Weight: 15,
	},
	{
		Key:    "deadline",
		Name:   "Prazo definido",
		Weight: 10,
	},
	{
		Key:    "public_data",
		Name:   "Dados públicos disponíveis",
		Weight: 15,
	},
	{
		Key:    "historical_baseline",
		Name:   "Histórico comparável",
		Weight: 15,
	},
	{
		Key:    "evidence_plausibility",
		Name:   "Plausibilidade baseada em evidências públicas",
		Weight: 20,
	},
	{
		Key:    "risks_dependencies",
		Name:   "Dependências e riscos",
		Weight: 15,
	},
}
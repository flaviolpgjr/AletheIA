package domain

var ScoringModelV1 = []Criterion{
	{
		Key:    "clarity",
		Name:   "Clareza da promessa",
		Weight: 15,
	},
	{
		Key:    "measurability",
		Name:   "Mensurabilidade",
		Weight: 20,
	},
	{
		Key:    "deadline",
		Name:   "Prazo definido",
		Weight: 10,
	},
	{
		Key:    "public_data",
		Name:   "Dados públicos disponíveis",
		Weight: 25,
	},
	{
		Key:    "historical_baseline",
		Name:   "Histórico comparável",
		Weight: 15,
	},
	{
		Key:    "risks_dependencies",
		Name:   "Dependências e riscos",
		Weight: 15,
	},
}
package domain

import "time"

type PublicDataBaseline struct {
	ID          int64
	Indicator   string
	Scope       string
	Value       float64
	Unit        string
	Source      string
	Reference   string
	CollectedAt time.Time
	CreatedAt   time.Time
}
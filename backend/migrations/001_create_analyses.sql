CREATE TABLE IF NOT EXISTS analyses (
    id UUID PRIMARY KEY,
    promise_text TEXT NOT NULL,
    score INTEGER NOT NULL,
    confidence INTEGER NOT NULL,
    analysis_data JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
ALTER TABLE analyses
ADD COLUMN IF NOT EXISTS promise_hash TEXT;

CREATE UNIQUE INDEX IF NOT EXISTS idx_analyses_promise_hash
ON analyses (promise_hash);
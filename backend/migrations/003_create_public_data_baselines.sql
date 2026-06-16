CREATE TABLE public_data_baselines (
    id BIGSERIAL PRIMARY KEY,

    indicator VARCHAR(100) NOT NULL,
    scope VARCHAR(50) NOT NULL,

    value NUMERIC NOT NULL,
    unit VARCHAR(50) NOT NULL,

    source VARCHAR(255) NOT NULL,
    reference VARCHAR(255) NOT NULL,

    collected_at TIMESTAMP NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_public_data_baselines_indicator_scope
ON public_data_baselines(indicator, scope);

CREATE INDEX idx_public_data_baselines_indicator
ON public_data_baselines(indicator);

CREATE INDEX idx_public_data_baselines_scope
ON public_data_baselines(scope);
-- +goose Up
CREATE TABLE games (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    version SERIAL,
    game_state JSONB,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- +goose Down
DROP TABLE games;

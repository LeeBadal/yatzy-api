-- +goose Up
CREATE TABLE games (
    id UUID NOT NULL,
    version SERIAL NOT NULL,
    game_state JSONB,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (id, version)
);

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION increment_version()
RETURNS TRIGGER AS $$
BEGIN
    NEW.version = (
        SELECT COALESCE(MAX(version), 0) + 1
        FROM games
        WHERE id = NEW.id
    );
    RETURN NEW;
END;
$$
LANGUAGE plpgsql;
-- +goose StatementEnd

-- Create a trigger to call the increment_version function before insert
CREATE TRIGGER increment_version_trigger
BEFORE INSERT ON games
FOR EACH ROW
EXECUTE FUNCTION increment_version();

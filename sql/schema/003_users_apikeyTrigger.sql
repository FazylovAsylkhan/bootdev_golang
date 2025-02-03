-- +goose Up
ALTER TABLE users ALTER COLUMN api_key SET NOT NULL;
--CREATE FUNCTION generate_api_key() RETURNS TRIGGER AS $$
--BEGIN
--    NEW.api_key := encode(sha256(random()::text::bytea), 'hex');
--  RETURN NEW;
--END;
--$$ LANGUAGE plpgsql;

--CREATE TRIGGER set_api_key
--BEFORE INSERT ON users
--FOR EACH ROW
--WHEN (NEW.api_key IS NULL)
--EXECUTE FUNCTION generate_api_key();

-- +goose Down
ALTER TABLE users ALTER COLUMN api_key DROP NOT NULL;
--DROP TRIGGER IF EXISTS set_api_key ON users;
--DROP FUNCTION IF EXISTS generate_api_key();
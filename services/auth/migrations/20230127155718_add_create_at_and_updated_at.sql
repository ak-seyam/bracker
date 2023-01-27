-- +goose Up
-- +goose StatementBegin
-- users
ALTER TABLE users 
ADD COLUMN created_at TIMESTAMP;

ALTER TABLE users 
ADD COLUMN updated_at TIMESTAMP;

-- groups
ALTER TABLE groups 
ADD COLUMN created_at TIMESTAMP;

ALTER TABLE groups 
ADD COLUMN updated_at TIMESTAMP;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- users
ALTER TABLE users
DROP COLUMN IF EXISTS create_at;

ALTER TABLE users
DROP COLUMN IF EXISTS updated_at;

-- groups
ALTER TABLE groups
DROP COLUMN create_at;

ALTER TABLE groups
DROP COLUMN updated_at;
-- +goose StatementEnd

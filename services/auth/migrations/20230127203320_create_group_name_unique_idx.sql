-- +goose Up
-- +goose StatementBegin
ALTER TABLE groups 
ADD CONSTRAINT unique_groups_name UNIQUE(name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE groups DROP CONSTRAINT unique_groups_name;
-- +goose StatementEnd

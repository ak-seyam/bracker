-- +goose Up
-- +goose StatementBegin
INSERT INTO groups(id, name, created_at, updated_at)
VALUES ('7c4f6839-420b-4f49-a015-1a47af2f915e', 'users', now(), now());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM groups 
WHERE name = 'users';
-- +goose StatementEnd

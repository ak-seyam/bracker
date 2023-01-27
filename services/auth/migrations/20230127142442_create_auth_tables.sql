-- +goose Up
-- +goose StatementBegin
CREATE TABLE groups(
	id   varchar(36) PRIMARY KEY,
	name varchar(255) NOT NULL
);

CREATE TABLE users(
	id       varchar(36) PRIMARY KEY,
	name     varchar(255) NOT NULL,
	username varchar(255) NOT NULL,
	password varchar(255) NOT NULL
);

ALTER TABLE users ADD CONSTRAINT unique_username
UNIQUE INCLUDE(username);

CREATE TABLE user_groups(
	user_id  varchar(36) NOT NULL,
	group_id varchar(36) NOT NULL
);

ALTER TABLE user_groups ADD CONSTRAINT fk_user_groups_user_id 
FOREIGN KEY(user_id) REFERENCES users(id);

ALTER TABLE user_groups ADD CONSTRAINT fk_user_groups_group_id
FOREIGN KEY(user_id) REFERENCES groups(id);

ALTER TABLE user_groups ADD CONSTRAINT pk_user_groups
PRIMARY KEY INCLUDE(user_id, group_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_groups;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS groups;
-- +goose StatementEnd

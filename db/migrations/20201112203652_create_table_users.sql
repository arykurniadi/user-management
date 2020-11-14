-- +goose Up
-- +goose StatementBegin
CREATE TABLE `users` (
	id varchar(50) NOT NULL,
    data text NOT NULL,
    role_id varchar(50) DEFAULT NULL,
	PRIMARY KEY (id),
	KEY users_id_index (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `users`;
-- +goose StatementEnd

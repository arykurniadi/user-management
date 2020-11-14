-- +goose Up
-- +goose StatementBegin
CREATE TABLE `roles` (
	id varchar(50) NOT NULL,
    data text NOT NULL,
	PRIMARY KEY (id),
	KEY roles_id_index (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `roles`;
-- +goose StatementEnd

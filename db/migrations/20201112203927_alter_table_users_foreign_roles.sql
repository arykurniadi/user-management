-- +goose Up
-- +goose StatementBegin
ALTER TABLE `users` ADD FOREIGN KEY(role_id) REFERENCES roles(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd

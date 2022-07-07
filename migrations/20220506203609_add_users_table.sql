-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id                SERIAL PRIMARY KEY,
    tg_username VARCHAR     NOT NULL UNIQUE,
    tg_chat_id  INTEGER     NOT NULL,
    created_at        TIMESTAMPTZ NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table users;
-- +goose StatementEnd

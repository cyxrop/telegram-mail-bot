-- +goose Up
-- +goose StatementBegin
CREATE TABLE mailboxes
(
    id              SERIAL PRIMARY KEY,
    mail            VARCHAR     NOT NULL UNIQUE,
    password        VARCHAR     NOT NULL,
    user_id         INTEGER     NOT NULL,
    last_message_id INTEGER     NOT NULL,
    polled_at       TIMESTAMPTZ NOT NULL,

    CONSTRAINT fk_mailbox_user_id
        FOREIGN KEY (user_id)
            REFERENCES users (id)
            ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP table mailboxes;
-- +goose StatementEnd

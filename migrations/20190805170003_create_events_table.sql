-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS events
(
    id             INT PRIMARY KEY,
    name           VARCHAR(255) NOT NULL,
    description    TEXT         NULL,
    location       VARCHAR(255) NOT NULL,
    availability   INT          NOT NULL,
    created_at     TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMP    NOT NULL DEFAULT NOW(),
    deleted_at     TIMESTAMP    NULL
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS events;
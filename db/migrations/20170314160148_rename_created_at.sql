
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE users RENAME COLUMN createdAt TO created_at;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back



-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE transactions ADD COLUMN completed_at timestamp with time zone;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back


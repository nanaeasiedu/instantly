
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS transactions (id serial, amount numeric, type varchar(255), mno varchar(255), reference text, message text, mobile_number varchar(255), receive_token text, network_id text, status text, reference_id text, created_at timestamp with time zone, updated_at timestamp with time zone);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE transactions;

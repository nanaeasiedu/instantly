
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE transactions DROP COLUMN wallet_id;
ALTER TABLE transactions ADD COLUMN user_id int references users (id);
ALTER TABLE users ADD COLUMN token varchar(255);
ALTER TABLE users ADD COLUMN callback_url text;
ALTER TABLE users ADD COLUMN current_balance numeric;
ALTER TABLE users ADD COLUMN mobile_number varchar(255);
ALTER TABLE users ADD COLUMN network_operator varchar(255);
DROP TABLE wallets;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back


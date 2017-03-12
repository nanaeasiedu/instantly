
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS users (id serial primary key, email_address varchar(255) unique not null, password varchar(255), password_hash text, createdAt timestamp with time zone, updated_at timestamp with time zone);
CREATE TABLE IF NOT EXISTS wallets (id serial primary key, uid varchar(255), token varchar(255), callback_url text, current_balance varchar(255), mobile_number varchar(255), network_operator varchar(255), created_at timestamp with time zone, updated_at timestamp with time zone, user_id int references users (id));
ALTER TABLE transactions ADD COLUMN wallet_id int references wallets (id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;
DROP TABLE wallets;


-- +migrate Up
alter table users alter column registered_at set default current_timestamp;
-- +migrate Down
alter table users alter column registered_at drop default;

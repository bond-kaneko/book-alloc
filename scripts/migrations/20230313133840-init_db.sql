-- +migrate Up
CREATE
EXTENSION if not exists pgcrypto;

drop table if exists users;
create table users
(
    id bigserial not null constraint users_pk primary key,
    auth0_id text not null,
    registered_at timestamp not null,
    created_at timestamp not null,
    updated_at timestamp
);

-- +migrate Down
drop table users;

-- +migrate Up
CREATE
EXTENSION if not exists pgcrypto;

drop table if exists users;
create table users
(
    id uuid not null DEFAULT gen_random_uuid() constraint users_pk primary key,
    auth0_id text not null,
    email text not null,
    name text not null,
    registered_at timestamp not null,
    created_at timestamp not null,
    updated_at timestamp
);

-- +migrate Down
drop table users;

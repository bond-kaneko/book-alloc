-- +migrate Up
CREATE EXTENSION if not exists pgcrypto;

drop table if exists users;
create table users
(
    id           uuid        not null
        default gen_random_uuid()
        constraint users_pk
            primary key,
    name         varchar(50) not null,
    email        text        not null
        constraint users_uk
            unique,
    password     text        not null,
    registered_at timestamp   not null,
    created_at   timestamp,
    updated_at   timestamp
);
create index users_email_name_index
    on users (email, name);

drop table if exists allocations;
create table allocations
(
    id         serial      not null
        constraint allocations_pk
            primary key,
    user_id    uuid      not null
        constraint allocations_users_fk
            references users,
    name       text      not null,
    share      integer   not null,
    created_at timestamp not null,
    updated_at timestamp not null
);

drop table if exists reading_histories;
create table reading_histories
(
    id            serial    not null
        constraint reading_histories_pk
            primary key,
    allocation_id serial
        constraint reading_histories_allocations_fk
            references allocations,
    isbn          varchar(13),
    title         text      not null,
    status        smallint  not null,
    times         integer   not null,
    start_at      timestamp not null,
    end_at        timestamp,
    rating        smallint,
    comment       text,
    created_at    timestamp not null,
    updated_at    timestamp not null
);
-- +migrate Down
drop table reading_histories;
drop table allocations;
drop table users;

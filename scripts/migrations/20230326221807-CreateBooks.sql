-- +migrate Up
drop table if exists books;
create table books
(
    id            bigserial not null,
    allocation_id bigint    not null,
    title         text      not null,
    status        int       not null,
    start_at      timestamp,
    end_at        timestamp,
    created_at    timestamp not null,
    updated_at    timestamp,
    constraint books_pk primary key (id),
    constraint books_allocations_fk foreign key (allocation_id) references allocations (id)
);
-- +migrate Down
drop table books;


-- +migrate Up
drop table if exists allocations;
create table allocations
(
    id bigserial not null,
    user_id uuid not null,
    name text not null,
    share int not null ,
    is_active bool not null ,
    created_at timestamp not null,
    updated_at timestamp,
    constraint allocations_pk primary key (id),
    constraint allocations_users_fk foreign key (user_id) references users(id)
);
-- +migrate Down
drop table allocations;

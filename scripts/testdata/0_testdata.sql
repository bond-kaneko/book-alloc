insert into users (auth0_id, email, name, registered_at, created_at, updated_at)
values ('DUMMY_ID', 'test@example.com', 'Test User', now(), now(), null);

insert into allocations (user_id, name, share, is_active, created_at, updated_at)
values ((select id from users order by id limit 1), 'one', 1, true, now(), null),
       ((select id from users order by id limit 1), 'two', 2, true, now(), null);

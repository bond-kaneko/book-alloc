insert into users (name, email, password, registered_at, created_at, updated_at)
values ('test', 'test@example.com', 'password', '2023-03-13 16:21:00', '2023-03-13 16:21:00', '2023-03-13 16:21:00');

insert into allocations (user_id, name, share, created_at, updated_at)
values ((select id from users order by users.created_at limit 1), '技術書', 1, '2023-03-13 16:21:00', '2023-03-13 16:21:00'),
       ((select id from users order by users.created_at limit 1), '専門外基礎', 1, '2023-03-13 16:21:00', '2023-03-13 16:21:00');

insert into reading_histories (allocation_id, isbn, title, status, times, start_at, end_at, rating, comment, created_at, updated_at)
values (1, null, 'Go言語プログラミングエッセンス', 1, 1, '2023-03-13 16:21:00', null, null, null, '2023-03-13 16:21:00', '2023-03-13 16:21:00'),
       (1, null, '論理哲学論考', 1, 1, '2023-03-13 16:21:00', null, null, null, '2023-03-13 16:21:00', '2023-03-13 16:21:00')

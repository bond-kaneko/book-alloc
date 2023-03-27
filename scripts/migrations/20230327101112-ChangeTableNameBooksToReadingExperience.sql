
-- +migrate Up
alter table books rename to reading_experiences;
-- +migrate Down
alter table reading_experiences rename to books;

#!/bin/bash

sql-migrate up -config=./config/sql-migrate/dbconfig.yml -env=local
sql-migrate up -config=./config/sql-migrate/dbconfig.yml -env=test

files="./scripts/testdata/*"
for f in $files; do
  PGPASSWORD=password psql -h db -U user book_alloc -e < "${f}"
  PGPASSWORD=password psql -h test_db -U user book_alloc_test -e < "${f}"
done

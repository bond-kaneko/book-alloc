#!/bin/bash

sql-migrate up -config=./config/sql-migrate/dbconfig.yml -env=local

files="./scripts/testdata/*"
for f in $files; do
  PGPASSWORD=password psql  -h db -U user book_alloc -e < "${f}"
done

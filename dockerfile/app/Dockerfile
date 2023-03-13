FROM golang:1.20.2

RUN apt update && apt install -y postgresql-client

RUN go install -v github.com/rubenv/sql-migrate/...@v1.1.1
RUN go install -v github.com/cosmtrek/air/...@v1.40.4
RUN go install -v github.com/lib/pq/...@v1.10.4

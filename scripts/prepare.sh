#!/bin/bash

DB_USER="validator"
DB_PASSWORD="val1dat0r"
DB_NAME="project-sem-1"
DB_TABLE_NAME="prices"

if ! command -v go &> /dev/null; then
  echo "Installing Go..."
  sudo apt update
  sudo apt install -y golang
fi

go mod init example.com/project_sem
go get -u github.com/lib/pq


if ! command -v psql &> /dev/null; then
  echo "Installing Postgres..."
  sudo apt update
  sudo apt install -y postgresql postgresql-contrib
fi

sudo systemctl start postgresql

sudo -u postgres psql << EOF
DO
\$do\$
BEGIN
   IF NOT EXISTS (
      SELECT
      FROM pg_catalog.pg_user
      WHERE usename = '$DB_USER') THEN
      CREATE USER $DB_USER WITH PASSWORD '$DB_PASSWORD';
   END IF;
END
\$do\$;

CREATE DATABASE $DB_NAME OWNER $DB_USER;
EOF

sudo -u postgres psql -d $DB_NAME << EOF
CREATE TABLE IF NOT EXISTS $DB_TABLE_NAME (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    category VARCHAR,
    price DECIMAL(10, 2),
    create_date TIMESTAMP
);
EOF

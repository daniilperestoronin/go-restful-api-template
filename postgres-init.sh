#!/bin/sh -e
psql --variable=ON_ERROR_STOP=1 --username "postgres" <<-EOSQL
    CREATE USER record_user WITH password 'record_user';
    CREATE DATABASE records;
    GRANT ALL PRIVILEGES ON DATABASE records TO record_user;
    CREATE TABLE record (
      id bigserial PRIMARY KEY,
      title VARCHAR(255) NOT NULL,
      text VARCHAR(255) NOT NULL
    );
EOSQL
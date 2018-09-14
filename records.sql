CREATE USER record_user WITH password 'record_user';

CREATE DATABASE records;

CREATE TABLE record (
  id bigserial PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  text VARCHAR(255) NOT NULL
);

GRANT ALL PRIVILEGES ON record TO record_user;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO record_user;

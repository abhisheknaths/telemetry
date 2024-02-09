CREATE USER myuser WITH PASSWORD 'mypass';
CREATE DATABASE telmetry;

GRANT ALL PRIVILEGES ON DATABASE telmetry TO myuser;

\c telmetry myuser;

CREATE SCHEMA telmetry_schema;

CREATE TABLE telmetry_schema.telmetry_users (user_id serial PRIMARY KEY, user_name TEXT);

INSERT INTO telmetry_schema.telmetry_users (user_name) VALUES ('Charles'), ('Hamilton'), ('Max');

CREATE TABLE telmetry_schema.telmetry_user_detail (user_id INTEGER, user_address TEXT);

INSERT INTO telmetry_schema.telmetry_user_detail (user_id, user_address) VALUES (1,'Monaco'), (2,'UK'), (3,'Amsterdam');

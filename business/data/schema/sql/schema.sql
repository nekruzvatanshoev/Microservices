CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name varchar,
    last_name varchar,
    email varchar NOT NULL UNIQUE ,
    date_created timestamp NOT NULL,
    status varchar  NOT NULL,
    password varchar NOT NULL
);

--
-- CREATE TABLE IF NOT EXISTS business (
--     id SERIAL PRIMARY KEY,
--     name varchar NOT NULL UNIQUE,
--     address varchar
-- );
--
--
-- CREATE TABLE IF NOT EXISTS userBusinesses (
--     id SERIAL PRIMARY KEY,
--     user_id int FOREIGN KEY REFERENCES users(id),
--     business_id int FOREIGN KEY REFERENCES
-- );
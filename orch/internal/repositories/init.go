package repository

var Build = `
CREATE TABLE IF NOT EXISTS users 
(
id text,
login text UNIQUE,
password text
);

CREATE TABLE IF NOT EXISTS expressions 
(
userID text,
id text PRIMARY KEY,
status text,
expression text,
result text
);`

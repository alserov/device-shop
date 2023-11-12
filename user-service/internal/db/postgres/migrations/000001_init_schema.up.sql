CREATE TABLE IF NOT EXISTS users
(
    UUID          text,
    Username      varchar(50),
    Password      varchar(100),
    Email         text,
    Cash          int CHECK (Cash > 0),
    Refresh_Token TEXT,
    Role          varchar(5),
    Created_At    timestamptz NOT NULL DEFAULT (now())
);

INSERT INTO users (UUID, Username, Password, Email, Cash, Refresh_Token, Role)
VALUES ('0', 'admin', 'qwerty', '', 0, '', 'admin');
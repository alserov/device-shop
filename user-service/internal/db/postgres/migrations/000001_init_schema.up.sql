CREATE TABLE IF NOT EXISTS users
(
    UUID          text,
    Username      varchar(50),
    Password      varchar(100),
    Email         text,
    Cash          float CHECK (Cash >= 0),
    Refresh_Token TEXT,
    Role          varchar(5),
    Created_At    timestamptz NOT NULL DEFAULT (now())
);

INSERT INTO users (UUID, Username, Password, Email, Cash, Refresh_Token, Role)
VALUES ('0', 'admin', '$2a$10$v6bGzITsBgUZjv/vCzrEVers/kP9sO3fPqup.wmaxCp3WEe6m1kqa', '', 0, '', 'admin');
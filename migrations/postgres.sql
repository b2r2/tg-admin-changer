CREATE TABLE IF NOT EXISTS users
(
   id                   serial              NOT NULL,
   first                varchar(255)        NOT NULL,
   username             varchar(255)        NOT NULL,
   created_at           timestamp           NOT NULL DEFAULT current_timestamp,
   updated_at           timestamp           NOT NULL DEFAULT current_timestamp,
   primary key(id,username)
);


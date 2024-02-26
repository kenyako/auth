-- +goose Up
-- create type user_role as Enum('USER', 'ADMIN');

create table users (
    id serial primary key,
    name text not null,
    email text,
    password text not null,
    password_confirm text not null,
    role text not null,
    created_at timestamp not null default clock_timestamp(),
    updated_at timestamp not null default clock_timestamp()
);

-- +goose Down
drop table users;

create table IF NOT EXISTS todos(
    id serial primary key,
    name VARCHAR(50) not null default(''),
    is_check boolean not null default(false),
    create_at TIMESTAMP default(now())
);
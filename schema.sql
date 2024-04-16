create table if not exists todos
(
    id uuid primary key,
    title varchar(50) not null,
    description varchar(1000) null,
    completed boolean not null,
    created_at timestamp not null,
    updated_at timestamp not null
);

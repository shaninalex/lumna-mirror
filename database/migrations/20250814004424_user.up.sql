CREATE TABLE public.users (
    id              UUID                primary key default uuid_generate_v4(),
    name            varchar(255),
    email           varchar(255)        unique not null,
    password_hash   varchar             not null,
    active          BOOLEAN             NOT NULL DEFAULT False,
    created_at      timestamp           default now(),
    updated_at      timestamp           default now()
);
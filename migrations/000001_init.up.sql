create table users(
    id serial not null unique,
    name varchar not null,
    surname varchar,
    username varchar,
    password varchar not null,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP NULL
);
create table authors(
    id serial not null unique,
    name varchar,
    surname varchar
);
create table books(
    id serial not null unique,
    name varchar not null,
    author_id integer not null references authors(id) ON DELETE CASCADE
);
create table book_history(
    id serial not null unique,
    client_id integer not null references users(id) ON DELETE CASCADE,
    book_id integer not null references books(id) ON DELETE CASCADE,
    taken_at timestamp,
    returned_at timestamp NULL
);

insert into authors(name, surname) values('Margulan', 'Seisembai');
insert into authors(name, surname) values('Deil', 'Karnegi');

insert into books(name, author_id) values('Winner', 1);
insert into books(name, author_id) values('Looser', 2   );

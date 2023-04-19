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

create table user_balance(
    user_id int primary key,
    balance float
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

create table book_prices(
    id serial not null unique,
    rental_price float CHECK ( rental_price >= 0 ),
    purchase_price float CHECK ( purchase_price >= 0 ),
    book_id integer not null references books(id) ON DELETE CASCADE
);

create table transactions(
    id serial not null unique,
    client_id integer not null references users(id) ON DELETE CASCADE,
    book_id integer not null references books(id) ON DELETE CASCADE,
    transaction_type varchar(20),
    amount float,
    taken_at timestamp,
    returned_at timestamp
);

create table statistics(
   book_id int primary key,
   total_income float,
   foreign key(book_id) references books(id)
);



insert into authors(name, surname) values('Margulan', 'Seisembai');
insert into authors(name, surname) values('Deil', 'Karnegi');

insert into books(name, author_id) values('Winner', 1);
insert into books(name, author_id) values('Looser', 2);

insert into book_prices(rental_price, purchase_price, book_id) values (200, 1000, 1);
insert into book_prices(rental_price, purchase_price, book_id) values (100, 2000, 2);


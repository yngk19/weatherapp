CREATE TABLE cities (
    id serial primary key,
    name varchar(60), 
    country varchar(60),
    lat float,
    lon float
);

CREATE TABLE forecasts (
    id serial primary key,
    temp float,
    city int references cities(id) on delete cascade,
    data jsonb not null
);

CREATE TABLE users (
    id serial primary key,
    username varchar(50),
    password varchar(100)
);
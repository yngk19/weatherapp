CREATE TABLE cities (
    name varchar(60) primary key, 
    country varchar(60),
    lat float,
    lon float
);

CREATE TABLE forecasts (
    id serial primary key,
    temp float,
    city varchar(60) references cities(name) on delete cascade,
    data jsonb not null
);

CREATE TABLE users (
    id serial primary key,
    username varchar(50),
    password varchar(100)
);
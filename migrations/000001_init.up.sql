CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50),
    password VARCHAR(100)
);

CREATE TABLE cities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(60) UNIQUE, 
    country VARCHAR(60),
    lat FLOAT,
    lon FLOAT,
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE forecasts (
    id SERIAL PRIMARY KEY,
    temp FLOAT,
    city_name varchar(60) UNIQUE,
    prediction_date timestamp,
    data JSONB NOT NULL,
    FOREIGN KEY (city_name) REFERENCES cities(name) ON DELETE CASCADE
);

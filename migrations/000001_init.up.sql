CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,  
    password TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP 
);

CREATE TABLE cities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL, 
    country VARCHAR(100) NOT NULL,
    state VARCHAR(100) NOT NULL,
    lat FLOAT NOT NULL,  
    lon FLOAT NOT NULL, 
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    UNIQUE (lat, lon)
);

CREATE TABLE forecasts (
    id SERIAL PRIMARY KEY UNIQUE,
    temp FLOAT NOT NULL,
    city_id INT NOT NULL,
    predict_date DATE NOT NULL,
    detail_info JSONB NOT NULL,
    FOREIGN KEY (city_id) REFERENCES cities(id) ON DELETE CASCADE,
    UNIQUE (city_id, predict_date)
);

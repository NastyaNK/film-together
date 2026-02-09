CREATE TABLE film
(
    id     SERIAL PRIMARY KEY,
    name   VARCHAR(255) NOT NULL,
    year   INT          NOT NULL,
    plot   TEXT,
    genre  VARCHAR(50),
    rating FLOAT,
    image  TEXT
);

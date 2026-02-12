CREATE TABLE room
(
    id      SERIAL PRIMARY KEY,
    name    VARCHAR(255) NOT NULL,
    film_id INTEGER      NOT NULL,
    public  BOOLEAN,
    CONSTRAINT fk_film
        FOREIGN KEY (film_id)
            REFERENCES film (id)
            ON DELETE CASCADE
);

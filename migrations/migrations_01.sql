DROP TABLE url;

CREATE TABLE url (
    id SERIAL PRIMARY KEY,
    long_url TEXT,
    short_url TEXT
)
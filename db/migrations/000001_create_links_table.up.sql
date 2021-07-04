CREATE TABLE links
(
    id                SERIAL PRIMARY KEY,
    created_datetime  TIMESTAMP    NOT NULL,
    url               VARCHAR(255) NOT NULL,
    code              VARCHAR(50)  NOT NULL,
    hits              INT          NOT NULL,
    last_hit_datetime TIMESTAMP DEFAULT NULL
);

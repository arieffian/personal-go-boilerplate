CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users(
    id VARCHAR(64) PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(512) NOT NULL
);
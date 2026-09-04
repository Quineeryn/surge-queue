CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255),

    created_at TIMESTAMP,
    updated_at TIMESTAMP 

)
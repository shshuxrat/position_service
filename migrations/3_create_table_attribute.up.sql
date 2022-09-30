CREATE TABLE IF NOT EXISTS attribute(
    id uuid PRIMARY KEY , 
    name VARCHAR(255) NOT NULL,
    "type" VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT(Now()),
	updated_at TIMESTAMP DEFAULT(Now())
);

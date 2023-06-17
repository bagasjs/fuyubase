CREATE TABLE users (
	id BIGINT NOT NULL,
	email VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	name VARCHAR(255) NOT NULL,
	
	CONSTRAINT UC_users UNIQUE(id,email),
	CONSTRAINT PK_users PRIMARY KEY(id)
)
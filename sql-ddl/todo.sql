CREATE TABLE todos (
	id serial PRIMARY KEY,
	action VARCHAR(100) NOT NULL,
	done boolean NOT NULL,
	username VARCHAR(50) NOT NULL
)
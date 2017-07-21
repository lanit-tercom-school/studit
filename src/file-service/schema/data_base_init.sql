CREATE TABLE file (
	id                SERIAL                                ,
	created           TIMESTAMP WITH TIME ZONE      NOT NULL,
	name              TEXT                          NOT NULL,
	path              TEXT                          NOT NULL,
	user_id           INT                           NOT NULL,

	CONSTRAINT file_pk PRIMARY KEY (id),
	CONSTRAINT file_fk0 FOREIGN KEY (user_id) REFERENCES "user"(id)
);
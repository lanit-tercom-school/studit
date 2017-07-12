CREATE TABLE file (
	id                SERIAL                        NOT NULL,
	user_id           INT                           NOT NULL,
	name              VARCHAR(100)                  NOT NULL,
	path              VARCHAR(255)                  NOT NULL,
	date_of_creation  TIMESTAMP WITH TIME ZONE      NOT NULL,

	CONSTRAINT file_pk PRIMARY KEY (id)
) WITH (
OIDS = FALSE
);	
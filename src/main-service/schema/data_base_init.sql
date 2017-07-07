CREATE TYPE status AS ENUM ('opened', 'started', 'ended');

CREATE TABLE project (
	id                SERIAL                                ,
	name              VARCHAR(100)                  NOT NULL,
	description       TEXT                          NOT NULL,
	date_of_creation  TIMESTAMP WITH TIME ZONE      NOT NULL,
	logo              VARCHAR(1000)                 NOT NULL,
	tags              TEXT[]                        NOT NULL,
	status            status                        NOT NULL,
	
	CONSTRAINT project_pk PRIMARY KEY (id)
)

CREATE TABLE "user" (
	id                SERIAL                                ,
	nickname          VARCHAR(100)                  NOT NULL,
	description       TEXT                          NOT NULL,
	avatar            VARCHAR(1000)                 NOT NULL,
	
	CONSTRAINT user_pk PRIMARY KEY (id)
)

CREATE TABLE project_user (
	id                SERIAL                                ,
	project_id        INT                           NOT NULL,
	user_id           INT                           NOT NULL,
	signed_date       TIMESTAMP WITH TIME ZONE      NOT NULL,
	progress          INT                           NOT NULL,
	
	UNIQUE (project_id, user_id),
	CONSTRAINT project_user_pk PRIMARY KEY (id),
	CONSTRAINT project_user_fk0 FOREIGN KEY (project_id) REFERENCES project(id),
	CONSTRAINT project_user_fk1 FOREIGN KEY (user_id) REFERENCES "user"(id)
)

CREATE TABLE project_enroll (
	id                SERIAL                                ,
	project_id        INT                           NOT NULL,
	user_id           INT                           NOT NULL,
	enrolling_message TEXT                          NOT NULL,
	time              TIMESTAMP WITH TIME ZONE      NOT NULL,

	CONSTRAINT project_user_application_pk PRIMARY KEY (id),
	CONSTRAINT project_user_application_fk0 FOREIGN KEY (project_id) REFERENCES project(id),
	CONSTRAINT project_user_application_fk1 FOREIGN KEY (user_id) REFERENCES "user"(id),
	CONSTRAINT unique_pair_of_columns_for_project_sign_up UNIQUE (project_id, user_id)
)


CREATE TABLE user_contact (
	id                SERIAL                                ,
	contact           VARCHAR(255)                  NOT NULL,
	contact_type      VARCHAR(255)                  NOT NULL,
	user_id           INT                           NOT NULL,
	
	CONSTRAINT user_contact_pk PRIMARY KEY (id),
--  CONSTRAINT user_contact_fk0 FOREIGN KEY (contact_type_id) REFERENCES contact_type(id),
	CONSTRAINT user_contact_fk1 FOREIGN KEY (user_id) REFERENCES "user"(id)
)

/* type is one of email, tel, vk, icq */
CREATE TABLE contact_type (
	id                SERIAL                                ,
	type              VARCHAR(100)                  NOT NULL,
	
	CONSTRAINT contact_type_pk PRIMARY KEY (id)
)

CREATE TABLE news (
	id                SERIAL                                ,
	title             VARCHAR(255)                  NOT NULL,
	description       TEXT                          NOT NULL,
	image             VARCHAR(255)                  NOT NULL,
	date_of_creation  TIMESTAMP WITH TIME ZONE      NOT NULL,
	last_edit         TIMESTAMP                     NOT NULL,
	tags              VARCHAR(1000)                 NOT NULL,
	
	CONSTRAINT news_pk PRIMARY KEY (id)
)

/* project tasks */
CREATE TABLE task (
	id                SERIAL                                ,
	title             VARCHAR(100)                  NOT NULL,
	description       VARCHAR(255)                  NOT NULL,
	numberOfTask      INT                           NOT NULL,
	tags              VARCHAR                       NOT NULL,
	priority          INT                           NOT NULL,
	project_id        INT                           NOT NULL,
	project_user_id   INT                           NOT NULL,
	
	CONSTRAINT task_pk PRIMARY KEY (id),
	CONSTRAINT task_fk0 FOREIGN KEY (project_id) REFERENCES project(id),
	CONSTRAINT task_fk2 FOREIGN KEY (project_user_id) REFERENCES project_user(id)
)

/* common tags for course, tasks, ... */
CREATE TABLE tag (
	id                SERIAL                                ,
	name              VARCHAR(25)                   NOT NULL,
	
	CONSTRAINT tag_pk PRIMARY KEY (id)
)

CREATE TABLE tasks_tags_table (
	id                SERIAL                                ,
	task_id           INT                           NOT NULL,
	tag_id            INT                           NOT NULL,
	
	CONSTRAINT tasks_tags_table_pk PRIMARY KEY (id),
	CONSTRAINT tasks_tags_table_fk0 FOREIGN KEY (task_id) REFERENCES task(id),
	CONSTRAINT tasks_tags_table_fk1 FOREIGN KEY (tag_id) REFERENCES tag(id)
)

/* course lesson */
CREATE TABLE lesson (
	id                SERIAL                                ,
	title             VARCHAR(255)                  NOT NULL,
	course_id         INT                           NOT NULL,
	description       TEXT                          NOT NULL,
	rating            INT                           NOT NULL,
	
	CONSTRAINT lesson_pk PRIMARY KEY (id),
	CONSTRAINT lesson_fk0 FOREIGN KEY (course_id) REFERENCES course(id)
)

/* lesson task */
CREATE TABLE test (
	id                SERIAL                                ,
	title             TEXT                          NOT NULL,
	lesson_id         INT                           NOT NULL,
	
	CONSTRAINT test_pk PRIMARY KEY (id),
	CONSTRAINT test_fk0 FOREIGN KEY (lesson_id) REFERENCES lesson(id)
)

CREATE TABLE task_for_test (
	id                SERIAL                                ,
	question          TEXT                          NOT NULL,
	test_id           INT                           NOT NULL,
	
	CONSTRAINT task_for_test_pk PRIMARY KEY (id),
	CONSTRAINT task_for_test_fk0 FOREIGN KEY (test_id) REFERENCES test(id)
)

/* test question */
CREATE TABLE variant (
	id                SERIAL                                ,
	text              TEXT                          NOT NULL,
	correct_answer    BOOL                          NOT NULL,
	task_for_test_id  INT                           NOT NULL,
	
	CONSTRAINT variant_pk PRIMARY KEY (id),
	CONSTRAINT variant_fk0 FOREIGN KEY (task_for_test_id) REFERENCES task_for_test(id)
)

/* lesson practise */
CREATE TABLE practise (
	id                SERIAL                                ,
	lesson_id         INT                           NOT NULL,
	description       TEXT                          NOT NULL,
	
	CONSTRAINT practise_pk PRIMARY KEY (id),
	CONSTRAINT practise_fk0 FOREIGN KEY (lesson_id) REFERENCES lesson(id)
)

/* lesson video */
CREATE TABLE video (
	id                SERIAL                                ,
	lesson_id         INT                           NOT NULL,
	link              TEXT                          NOT NULL,
	
	CONSTRAINT video_pk PRIMARY KEY (id),
	CONSTRAINT video_fk0 FOREIGN KEY (lesson_id) REFERENCES lesson(id)
)

CREATE TABLE user_course (
	id                SERIAL                                ,
	user_id           INT                           NOT NULL,
	course_id         INT                           NOT NULL,
	date              TIMESTAMP WITH TIME ZONE      NOT NULL,
	progress          INT                           NOT NULL,
	
	CONSTRAINT user_course_pk PRIMARY KEY (id),
	CONSTRAINT user_course_fk0 FOREIGN KEY (course_id) REFERENCES course(id),
	CONSTRAINT user_course_fk1 FOREIGN KEY (user_id) REFERENCES "user"(id)
)

CREATE TABLE user_comments (
	id                SERIAL                                ,
	user_id           INT                           NOT NULL,
	-- в случае цитирования или комментирования комментария
	comment_id        INT                           NOT NULL,
	date              TIMESTAMP WITH TIME ZONE      NOT NULL,
	
	CONSTRAINT user_comments_pk PRIMARY KEY (id),
	CONSTRAINT user_comments_fk0 FOREIGN KEY (comment_id) REFERENCES comment(id),
	CONSTRAINT user_comments_fk1 FOREIGN KEY (user_id) REFERENCES "user"(id)
)

CREATE TABLE comment (
	id                SERIAL                                ,
	text              TEXT                          NOT NULL,
	
	CONSTRAINT comment_pk PRIMARY KEY (id)
)

CREATE TABLE course (
	id                SERIAL                        NOT NULL,
	title             VARCHAR(255)                  NOT NULL,
	description       TEXT                          NOT NULL,
	logo              VARCHAR(1000)                 NOT NULL,
	rating            REAL                          NOT NULL,
	
	CONSTRAINT course_pk PRIMARY KEY (id)
)

/* recommended courses for this course */
CREATE TABLE recommended_courses (
	id                SERIAL                                ,
	course_id         INT                           NOT NULL,
	link              VARCHAR(255)                  NOT NULL,
	
	CONSTRAINT recommended_courses_pk PRIMARY KEY (id),
	CONSTRAINT recommended_courses_fk0 FOREIGN KEY (course_id) REFERENCES course(id)
)

CREATE TABLE course_statistic (
	id                SERIAL                                ,
	hours             INT                           NOT NULL,
	course_id         INT                           NOT NULL,
	
	CONSTRAINT course_statistic_pk PRIMARY KEY (id),
	CONSTRAINT course_statistic_fk0 FOREIGN KEY (course_id) REFERENCES course(id)
)

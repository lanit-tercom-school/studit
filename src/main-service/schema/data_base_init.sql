-- table order convention - new entities as late as possible
CREATE TYPE status AS ENUM ('opened', 'started', 'ended');

<<<<<<< HEAD
CREATE TABLE project (
	id                SERIAL                                ,
	created           TIMESTAMP                     NOT NULL,
	description       TEXT                          NOT NULL,
	logo              TEXT                          NOT NULL,
	name              TEXT                          NOT NULL,
	status            status                        NOT NULL,
	tags              TEXT[]           NOT NULL DEFAULT '{}',
=======
/*Проект*/
CREATE TABLE "project" (
	"id"               SERIAL                   NOT NULL,
	"name"             VARCHAR(100)             NOT NULL,
	"description"      TEXT                     NOT NULL,
	"date_of_creation" TIMESTAMP WITH TIME ZONE NOT NULL,
	"logo"             VARCHAR(1000)            NOT NULL,
	"tags"             VARCHAR(1000)            NOT NULL,
	"status"           STATE                    NOT NULL,
	CONSTRAINT project_pk PRIMARY KEY ("id")
) WITH (
OIDS = FALSE
);

/*Пользователь*/
CREATE TABLE "user" (
	"id"          SERIAL        NOT NULL,
	"nickname"    VARCHAR(100)  NOT NULL,
	"description" TEXT          NOT NULL,
	"avatar"      VARCHAR(1000) NOT NULL,
	CONSTRAINT user_pk PRIMARY KEY ("id")
) WITH (
OIDS = FALSE
);


/*Связь пользователя и проекта, в котором участвует пользователь*/
CREATE TABLE "project_user" (
	"id"          SERIAL                   NOT NULL,
	"project_id"  BIGINT                   NOT NULL,
	"user_id"     BIGINT                   NOT NULL,
	"signed_date" TIMESTAMP WITH TIME ZONE NOT NULL,
	"progress"    INT                      NOT NULL,
	UNIQUE ("project_id", "user_id"),
	CONSTRAINT project_user_pk PRIMARY KEY ("id")
) WITH (
OIDS = FALSE
);



/*Связь пользователя и проекта, на который пользователь записан*/
CREATE TABLE "project_enroll" (
	"id"                SERIAL                   NOT NULL,
	"project_id"        BIGINT                   NOT NULL,
	"user_id"           BIGINT                   NOT NULL,
	"enrolling_message" TEXT                     NOT NULL, /*Сообщение для мастеров проекта, небольшое сопроводительное письмо*/
	"time"              TIMESTAMP WITH TIME ZONE NOT NULL, /*Дата, когда была подана заявка*/

	CONSTRAINT project_user_application_pk PRIMARY KEY ("id"),
	CONSTRAINT project_user_application_fk0 FOREIGN KEY ("project_id") REFERENCES "project" ("id"),
	CONSTRAINT project_user_application_fk1 FOREIGN KEY ("user_id") REFERENCES "user" ("id"),
	CONSTRAINT unique_pair_of_columns_for_project_sign_up UNIQUE ("project_id", "user_id")
) WITH (
OIDS = FALSE
);


/*Запись контакта пользователя*/
CREATE TABLE "user_contact" (
	"id"           SERIAL       NOT NULL,
	"contact"      VARCHAR(255) NOT NULL,
	"contact_type" VARCHAR(255) NOT NULL,
	"user_id"      BIGINT       NOT NULL,
	CONSTRAINT user_contact_pk PRIMARY KEY ("id")
) WITH (
OIDS = FALSE
);
CREATE TABLE file (
	id                SERIAL                        NOT NULL,
	user_id           INT                           NOT NULL,
	name              VARCHAR(100)                  NOT NULL,
	path              VARCHAR(255)                  NOT NULL,
	date_of_creation  TIMESTAMP WITH TIME ZONE      NOT NULL,

	CONSTRAINT file_pk PRIMARY KEY (id),
	CONSTRAINT file_fk0 FOREIGN KEY (user_id) REFERENCES "user"(id)
) WITH (
OIDS = FALSE
);

/*Тип контакта: телефон, мыло, vk, одноклассники и т.д.*/
CREATE TABLE "contact_type" (
	"id"   SERIAL       NOT NULL,
	"type" VARCHAR(100) NOT NULL,
	CONSTRAINT contact_type_pk PRIMARY KEY ("id")
) WITH (
OIDS = FALSE
);

/*Новость*/
CREATE TABLE "news" (
	"id"               SERIAL                   NOT NULL,
	"title"            VARCHAR(255)             NOT NULL,
	"description"      TEXT                     NOT NULL,
	"image"            VARCHAR(255)             NOT NULL,
	"date_of_creation" TIMESTAMP WITH TIME ZONE NOT NULL,
	"last_edit"        TIMESTAMP                NOT NULL,
	"tags"             VARCHAR(1000)            NOT NULL,
	CONSTRAINT news_pk PRIMARY KEY ("id")
) WITH (
OIDS = FALSE
);
>>>>>>> 22e6a1c18a1bf3a965b6ce8f633aac90cc542712

	CONSTRAINT project_pk PRIMARY KEY (id)
);

CREATE TABLE "user" (
	id                SERIAL                                ,
	nickname          TEXT                          NOT NULL,
	description       TEXT                          NOT NULL,
	avatar            TEXT                          NOT NULL,

	CONSTRAINT user_pk PRIMARY KEY (id)
);

CREATE TABLE project_user (
	id                SERIAL                                ,
	project_id        INT                           NOT NULL,
	user_id           INT                           NOT NULL,
	signed_date       TIMESTAMP                     NOT NULL,
	progress          INT                           NOT NULL,

	CONSTRAINT project_user_pk PRIMARY KEY (id),
	CONSTRAINT project_user_fk0 FOREIGN KEY (project_id) REFERENCES project(id),
	CONSTRAINT project_user_fk1 FOREIGN KEY (user_id) REFERENCES "user"(id),
	CONSTRAINT project_user_uq UNIQUE (project_id, user_id)
);

CREATE TABLE project_enroll (
	id                SERIAL                                ,
	project_id        INT                           NOT NULL,
	user_id           INT                           NOT NULL,
	enrolling_message TEXT                          NOT NULL,
	time              TIMESTAMP                     NOT NULL,

	CONSTRAINT project_enroll_pk PRIMARY KEY (id),
	CONSTRAINT project_enroll_fk0 FOREIGN KEY (project_id) REFERENCES project(id),
	CONSTRAINT project_enroll_fk1 FOREIGN KEY (user_id) REFERENCES "user"(id),
	CONSTRAINT project_enroll_uq UNIQUE (project_id, user_id)
);

-- worth to remove
-- type is one of email, tel, vk, icq
CREATE TABLE contact_type (
	id                SERIAL                                ,
	type              TEXT                          NOT NULL,
	
	CONSTRAINT contact_type_pk PRIMARY KEY (id)
);

CREATE TABLE user_contact (
	id                SERIAL                                ,
	contact           TEXT                          NOT NULL,
	contact_type_id   INT                           NOT NULL,
	user_id           INT                           NOT NULL,
	
	CONSTRAINT user_contact_pk PRIMARY KEY (id),
	CONSTRAINT user_contact_fk0 FOREIGN KEY (contact_type_id) REFERENCES contact_type(id),
	CONSTRAINT user_contact_fk1 FOREIGN KEY (user_id) REFERENCES "user"(id)
);

/* project tasks */
CREATE TABLE task (
	id                SERIAL                                ,
	title             TEXT                          NOT NULL,
	description       TEXT                          NOT NULL,
	numberOfTask      INT                           NOT NULL,
	tags              TEXT[]           NOT NULL DEFAULT '{}',
	priority          INT                           NOT NULL,
	project_id        INT                           NOT NULL,
	project_user_id   INT                           NOT NULL,
	
	CONSTRAINT task_pk PRIMARY KEY (id),
	CONSTRAINT task_fk0 FOREIGN KEY (project_id) REFERENCES project(id),
	CONSTRAINT task_fk1 FOREIGN KEY (project_user_id) REFERENCES project_user(id)
);

/* common tags for course, tasks, ... */
-- CREATE TABLE tag (
	-- id                SERIAL                                ,
	-- name              VARCHAR(25)                   NOT NULL,
	
	-- CONSTRAINT tag_pk PRIMARY KEY (id)
-- );

-- CREATE TABLE tasks_tags_table (
	-- id                SERIAL                                ,
	-- task_id           INT                           NOT NULL,
	-- tag_id            INT                           NOT NULL,
	
	-- CONSTRAINT tasks_tags_table_pk PRIMARY KEY (id),
	-- CONSTRAINT tasks_tags_table_fk0 FOREIGN KEY (task_id) REFERENCES task(id),
	-- CONSTRAINT tasks_tags_table_fk1 FOREIGN KEY (tag_id) REFERENCES tag(id)
-- );

CREATE TABLE comment (
	id                SERIAL                                ,
	text              TEXT                          NOT NULL,
	
	CONSTRAINT comment_pk PRIMARY KEY (id)
);

CREATE TABLE user_comment (
	id                SERIAL                                ,
	-- в случае цитирования или комментирования комментария
	comment_id        INT                           NOT NULL,
	date              TIMESTAMP                     NOT NULL,
	user_id           INT                           NOT NULL,
	
	CONSTRAINT user_comment_pk PRIMARY KEY (id),
	CONSTRAINT user_comment_fk0 FOREIGN KEY (comment_id) REFERENCES comment(id),
	CONSTRAINT user_comment_fk1 FOREIGN KEY (user_id) REFERENCES "user"(id)
);

CREATE TABLE course (
	id                SERIAL                                ,
	description       TEXT                          NOT NULL,
	logo              TEXT                          NOT NULL,
	rating            REAL                          NOT NULL,
	title             TEXT                          NOT NULL,
	
	CONSTRAINT course_pk PRIMARY KEY (id)
);

CREATE TABLE user_course (
	id                SERIAL                                ,
	course_id         INT                           NOT NULL,
	date              TIMESTAMP                     NOT NULL,
	progress          INT                           NOT NULL,
	user_id           INT                           NOT NULL,
	
	CONSTRAINT user_course_pk PRIMARY KEY (id),
	CONSTRAINT user_course_fk0 FOREIGN KEY (course_id) REFERENCES course(id),
	CONSTRAINT user_course_fk1 FOREIGN KEY (user_id) REFERENCES "user"(id)
);

/* course statistics */
CREATE TABLE statistics (
	id                SERIAL                                ,
	course_id         INT                           NOT NULL,
	hours             INT                           NOT NULL,
	
	CONSTRAINT statistics_pk PRIMARY KEY (id),
	CONSTRAINT statistics_fk0 FOREIGN KEY (course_id) REFERENCES course(id)
);

/* recommended course for other course */
CREATE TABLE recommended_course (
	id                SERIAL                                ,
	course_id         INT                           NOT NULL,
	link              TEXT                          NOT NULL,
	
	CONSTRAINT recommended_course_pk PRIMARY KEY (id),
	CONSTRAINT recommended_course_fk0 FOREIGN KEY (course_id) REFERENCES course(id)
);

/* course lesson */
CREATE TABLE lesson (
	id                SERIAL                                ,
	course_id         INT                           NOT NULL,
	description       TEXT                          NOT NULL,
	rating            INT                           NOT NULL,
	title             TEXT                          NOT NULL,
	
	CONSTRAINT lesson_pk PRIMARY KEY (id),
	CONSTRAINT lesson_fk0 FOREIGN KEY (course_id) REFERENCES course(id)
);

/* lesson practise */
CREATE TABLE practise (
	id                SERIAL                                ,
	description       TEXT                          NOT NULL,
	lesson_id         INT                           NOT NULL,
	
	CONSTRAINT practise_pk PRIMARY KEY (id),
	CONSTRAINT practise_fk0 FOREIGN KEY (lesson_id) REFERENCES lesson(id)
);

/* lesson video */
CREATE TABLE video (
	id                SERIAL                                ,
	lesson_id         INT                           NOT NULL,
	link              TEXT                          NOT NULL,
	
	CONSTRAINT video_pk PRIMARY KEY (id),
	CONSTRAINT video_fk0 FOREIGN KEY (lesson_id) REFERENCES lesson(id)
);

/* lesson task */
CREATE TABLE test (
	id                SERIAL                                ,
	lesson_id         INT                           NOT NULL,
	title             TEXT                          NOT NULL,
	
	CONSTRAINT test_pk PRIMARY KEY (id),
	CONSTRAINT test_fk0 FOREIGN KEY (lesson_id) REFERENCES lesson(id)
);

CREATE TABLE tests_task (
	id                SERIAL                                ,
	question          TEXT                          NOT NULL,
	test_id           INT                           NOT NULL,
	
	CONSTRAINT tests_task_pk PRIMARY KEY (id),
	CONSTRAINT tests_task_fk0 FOREIGN KEY (test_id) REFERENCES test(id)
);

/* test question */
CREATE TABLE variant (
	id                SERIAL                                ,
	correct_answer    BOOL                          NOT NULL,
	tests_task_id     INT                           NOT NULL,
	text              TEXT                          NOT NULL,
	
	CONSTRAINT variant_pk PRIMARY KEY (id),
	CONSTRAINT variant_fk0 FOREIGN KEY (tests_task_id) REFERENCES tests_task(id)
);

CREATE TABLE news (
	id                SERIAL                                ,
	created           TIMESTAMP                     NOT NULL,
	description       TEXT                          NOT NULL,
	image             TEXT                          NOT NULL,
	edited            TIMESTAMP                     NOT NULL,
	tags              TEXT[]           NOT NULL DEFAULT '{}',
	title             TEXT                          NOT NULL,
	
	CONSTRAINT news_pk PRIMARY KEY (id)
);

CREATE TABLE file (
	id                SERIAL                                ,
	created           TIMESTAMP                     NOT NULL,
	name              TEXT                          NOT NULL,
	path              TEXT                          NOT NULL,
	user_id           INT                           NOT NULL,

	CONSTRAINT file_pk PRIMARY KEY (id),
	CONSTRAINT file_fk0 FOREIGN KEY (user_id) REFERENCES "user"(id)
);

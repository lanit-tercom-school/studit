/*
0 - проект еще не начался, идет набор, и т.д.
1 - проект начался, ведутся лекции, разработка
2 - проект завершен, активность закончена
*/
/*TODO: refactor this enum*/
CREATE TYPE state AS ENUM ('0', '1', '2');

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

/*Связь мастеров и проектов*/
CREATE TABLE "project_master" (
	"id"          SERIAL                   NOT NULL,
	"project_id"  BIGINT                   NOT NULL,
	"master_id"   BIGINT                   NOT NULL,
	"signed_date" TIMESTAMP WITH TIME ZONE NOT NULL,

	UNIQUE ("project_id", "master_id"),
	CONSTRAINT project_master_pk PRIMARY KEY ("id"),
	CONSTRAINT project_master_fk0 FOREIGN KEY ("project_id") REFERENCES "project" ("id"),
	CONSTRAINT project_master_fk1 FOREIGN KEY ("master_id") REFERENCES "user" ("id")
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

/*Задача - часть проекта*/
CREATE TABLE "task" (
	"id"              SERIAL       NOT NULL,
	"title"           VARCHAR(100) NOT NULL,
	"description"     VARCHAR(255) NOT NULL,
	"numberOfTask"    INT          NOT NULL,
	"tags"            VARCHAR      NOT NULL,
	"priority"        INT          NOT NULL,
	"project_id"      BIGINT       NOT NULL,
	"project_user_id" BIGINT       NOT NULL,
	CONSTRAINT task_pk PRIMARY KEY ("id")
) WITH (
OIDS = FALSE
);

/*Общие теги (в данный момент эти теги общие для всех частей(модулей), т.е. у задач, курсов и т.д. одни и теже)*/
CREATE TABLE "tag" (
	"id"   SERIAL      NOT NULL,
	"name" VARCHAR(25) NOT NULL,
	CONSTRAINT tag_pk PRIMARY KEY ("id")
) WITH (
OIDS = FALSE
);


/*Связь задач и тегов*/
CREATE TABLE "tasks_tags_table" (
	"id"      SERIAL NOT NULL,
	"task_id" BIGINT NOT NULL,
	"tag_id"  BIGINT NOT NULL,
	CONSTRAINT tasks_tags_table_pk PRIMARY KEY ("id")
) WITH (
OIDS = FALSE
);

/*Урок - часть Курса*/
CREATE TABLE "lesson" (
	"id"          SERIAL       NOT NULL,
	"title"       VARCHAR(255) NOT NULL,
	"course_id"   BIGINT       NOT NULL,
	"description" TEXT         NOT NULL,
	"rating"      INT          NOT NULL,
	CONSTRAINT lesson_pk PRIMARY KEY ("id")
) WITH (
OIDS = FALSE
);

/*Тестовая часть урока*/
CREATE TABLE "test" (
	"id"        SERIAL NOT NULL,
	"title"     TEXT   NOT NULL,
	"lesson_id" BIGINT NOT NULL,
	CONSTRAINT test_pk PRIMARY KEY ("id")
) WITH (
OIDS = FALSE
);


/*Задание в тесте*/
CREATE TABLE "task_for_test" (
	"id"       SERIAL NOT NULL,
	"question" TEXT   NOT NULL,
	"test_id"  BIGINT NOT NULL,
	CONSTRAINT task_for_test_pk PRIMARY KEY ("id")
) WITH (
OIDS = FALSE
);


/*Вариант ответа в задании теста*/
CREATE TABLE "variant" (
	"id"               SERIAL NOT NULL,
	"text"             TEXT   NOT NULL,
	"correct_answer"   BOOL   NOT NULL,
	"task_for_test_id" BIGINT NOT NULL,
	CONSTRAINT variant_pk PRIMARY KEY ("id")
) WITH (
OIDS = FALSE
);


/*Практическая часть урока*/
CREATE TABLE "practise" (
	"id"          SERIAL NOT NULL,
	"lesson_id"   BIGINT NOT NULL,
	"description" TEXT   NOT NULL,
	CONSTRAINT practise_pk PRIMARY KEY ("id")
) WITH (
OIDS = FALSE
);


/*Видео(теоретическая) часть урока*/
CREATE TABLE "video" (
	"id"        SERIAL NOT NULL,
	"lesson_id" BIGINT NOT NULL,
	"link"      TEXT   NOT NULL,
	CONSTRAINT video_pk PRIMARY KEY ("id")
) WITH (
OIDS = FALSE
);


/*Связь пользователя и курса, на который он записан*/
CREATE TABLE "user_course" (
	"id"        SERIAL                   NOT NULL,
	"user_id"   BIGINT                   NOT NULL,
	"course_id" BIGINT                   NOT NULL,
	"date"      TIMESTAMP WITH TIME ZONE NOT NULL,
	"progress"  INT                      NOT NULL,
	CONSTRAINT user_course_pk PRIMARY KEY ("id")
) WITH (
OIDS = FALSE
);


/*Вспомогательная таблица, устанавливающая соответствие между комментарием и его автором*/
CREATE TABLE "user_comments" (
	"id"         SERIAL                   NOT NULL,
	"user_id"    BIGINT                   NOT NULL,
	"comment_id" BIGINT                   NOT NULL, -- в случае цитирования или комментирования комментария
	"date"       TIMESTAMP WITH TIME ZONE NOT NULL,
	CONSTRAINT user_comments_pk PRIMARY KEY ("id")
) WITH (
OIDS = FALSE
);

/*Комментарий пользователя*/
CREATE TABLE "comment" (
	"id"   SERIAL NOT NULL,
	"text" TEXT   NOT NULL,
	CONSTRAINT comment_pk PRIMARY KEY ("id")
) WITH (
OIDS = FALSE
);

/*Курс*/
CREATE TABLE "course" (
	"id"          SERIAL        NOT NULL,
	"title"       VARCHAR(255)  NOT NULL,
	"description" TEXT          NOT NULL,
	"logo"        VARCHAR(1000) NOT NULL,
	"rating"      REAL          NOT NULL,
	CONSTRAINT course_pk PRIMARY KEY ("id")
) WITH (
OIDS = FALSE
);


/*Список рекомендованных курсов для данного курса*/
CREATE TABLE "recomend_courses" (
	"id"        SERIAL       NOT NULL,
	"course_id" BIGINT       NOT NULL,
	"link"      VARCHAR(255) NOT NULL,
	CONSTRAINT recomend_courses_pk PRIMARY KEY ("id")
) WITH (
OIDS = FALSE
);


/*Дополнительные данные по курсу*/
CREATE TABLE "statistic" (
	"id"        SERIAL NOT NULL,
	"hours"     BIGINT NOT NULL,
	"course_id" BIGINT NOT NULL,
	CONSTRAINT statistic_pk PRIMARY KEY ("id")
) WITH (
OIDS = FALSE
);




ALTER TABLE "recomend_courses" ADD CONSTRAINT "recomend_courses_fk0" FOREIGN KEY ("course_id") REFERENCES "course"("id");

ALTER TABLE "statistic" ADD CONSTRAINT "statistic_fk0" FOREIGN KEY ("course_id") REFERENCES "course"("id");



ALTER TABLE "test" ADD CONSTRAINT "test_fk0" FOREIGN KEY ("lesson_id") REFERENCES "lesson"("id");

ALTER TABLE "task_for_test" ADD CONSTRAINT "task_for_test_fk0" FOREIGN KEY ("test_id") REFERENCES "test"("id");

ALTER TABLE "variant" ADD CONSTRAINT "variant_fk0" FOREIGN KEY ("task_for_test_id") REFERENCES "task_for_test"("id");

ALTER TABLE "practise" ADD CONSTRAINT "practise_fk0" FOREIGN KEY ("lesson_id") REFERENCES "lesson"("id");

ALTER TABLE "video" ADD CONSTRAINT "video_fk0" FOREIGN KEY ("lesson_id") REFERENCES "lesson"("id");


ALTER TABLE "task" ADD CONSTRAINT "task_fk0" FOREIGN KEY ("project_id") REFERENCES "project"("id");
ALTER TABLE "task" ADD CONSTRAINT "task_fk2" FOREIGN KEY ("project_user_id") REFERENCES "project_user"("id");


ALTER TABLE "tasks_tags_table" ADD CONSTRAINT "tasks_tags_table_fk0" FOREIGN KEY ("task_id") REFERENCES "task"("id");
ALTER TABLE "tasks_tags_table" ADD CONSTRAINT "tasks_tags_table_fk1" FOREIGN KEY ("tag_id") REFERENCES "tag"("id");

ALTER TABLE "lesson" ADD CONSTRAINT "lesson_fk0" FOREIGN KEY ("course_id") REFERENCES "course"("id");


ALTER TABLE "project_user" ADD CONSTRAINT "project_user_fk0" FOREIGN KEY ("project_id") REFERENCES "project"("id");
ALTER TABLE "project_user" ADD CONSTRAINT "project_user_fk1" FOREIGN KEY ("user_id") REFERENCES "user"("id");
/*
ALTER TABLE "user_contact" ADD CONSTRAINT "user_contact_fk0" FOREIGN KEY ("contact_type_id") REFERENCES "contact_type"("id");*/
ALTER TABLE "user_contact" ADD CONSTRAINT "user_contact_fk1" FOREIGN KEY ("user_id") REFERENCES "user"("id");

ALTER TABLE "user_course" ADD CONSTRAINT "user_course_fk0" FOREIGN KEY ("course_id") REFERENCES "course"("id");
ALTER TABLE "user_course" ADD CONSTRAINT "user_course_fk1" FOREIGN KEY ("user_id") REFERENCES "user"("id");

ALTER TABLE "user_comments" ADD CONSTRAINT "user_comments_fk0" FOREIGN KEY ("comment_id") REFERENCES "comment"("id");
ALTER TABLE "user_comments" ADD CONSTRAINT "user_comments_fk1" FOREIGN KEY ("user_id") REFERENCES "user"("id");

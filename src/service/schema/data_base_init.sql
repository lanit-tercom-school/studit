
CREATE TYPE state AS ENUM ('еще не начат', 'начат', 'завершен');

/*Проект*/
CREATE TABLE "project" (
	"id" serial NOT NULL,
	"name" varchar(100) NOT NULL,
	"description" TEXT NOT NULL,
	"date_of_creation" TIMESTAMP WITH TIME ZONE NOT NULL,
	"logo" varchar(1000) NOT NULL,
	"tags" varchar(1000) NOT NULL,
	"status" state NOT NULL,
	CONSTRAINT project_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

/*Пользователь*/
CREATE TABLE "user" (
  "id" serial NOT NULL,
  "login" varchar(100) NOT NULL,
  "password" varchar(100) NOT NULL,
  "nickname" varchar(100) NOT NULL,
  "description" TEXT NOT NULL,
  "avatar" varchar(1000) NOT NULL,
  "permission_level" INT NOT NULL,
  CONSTRAINT user_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);


/*Связь пользователя и проекта, в котором участвует пользователь*/
CREATE TABLE "project_user" (
  "id" serial NOT NULL,
  "project_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "signed_date" TIMESTAMP WITH TIME ZONE NOT NULL,
  "progress" int NOT NULL,
  UNIQUE ("project_id", "user_id"),
  CONSTRAINT project_user_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

/*Связь мастеров и проектов*/
CREATE TABLE "project_master" (
	"id" serial NOT NULL,
	"project_id" bigint NOT NULL,
	"master_id" bigint NOT NULL,
	"signed_date" TIMESTAMP WITH TIME ZONE NOT NULL,

	UNIQUE ("project_id", "master_id"),
	CONSTRAINT project_master_pk PRIMARY KEY ("id"),
	CONSTRAINT project_master_fk0 FOREIGN KEY ("project_id") REFERENCES "project"("id"),
	CONSTRAINT project_master_fk1 FOREIGN KEY ("master_id") REFERENCES "user"("id")
) WITH (
OIDS=FALSE
);

/*Связь пользователя и проекта, на который пользователь записан*/
CREATE TABLE "project_enroll" (
  "id" serial NOT NULL,
  "project_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,

  CONSTRAINT project_user_application_pk PRIMARY KEY ("id"),
  CONSTRAINT project_user_application_fk0 FOREIGN KEY ("project_id") REFERENCES "project"("id"),
  CONSTRAINT project_user_application_fk1 FOREIGN KEY ("user_id") REFERENCES "user"("id"),
  CONSTRAINT unique_pair_of_columns_for_project_sign_up UNIQUE ("project_id", "user_id")
) WITH (
  OIDS=FALSE
);


/*Запись контакта пользователя*/
CREATE TABLE "user_contact" (
	"id" serial NOT NULL,
	"contact" varchar(255) NOT NULL,
	"contact_type" varchar(255) NOT NULL,
	"user_id" bigint NOT NULL,
	CONSTRAINT user_contact_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);


/*Тип контакта: телефон, мыло, vk, одноклассники и т.д.*/
CREATE TABLE "contact_type" (
	"id" serial NOT NULL,
	"type" varchar(100) NOT NULL,
	CONSTRAINT contact_type_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

/*Новость*/
CREATE TABLE "news" (
	"id" serial NOT NULL,
	"title" varchar(255) NOT NULL,
	"description" TEXT NOT NULL,
	"image" varchar(255) NOT NULL,
	"date_of_creation" TIMESTAMP WITH TIME ZONE NOT NULL,
	"last_edit" TIMESTAMP NOT NULL,
	"tags" varchar(1000) NOT NULL,
	CONSTRAINT news_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

/*Задача - часть проекта*/
CREATE TABLE "task" (
	"id" serial NOT NULL,
	"title" varchar(100) NOT NULL,
	"description" varchar(255) NOT NULL,
	"numberOfTask" int NOT NULL,
	"tags" varchar NOT NULL,
	"priority" int NOT NULL,
	"project_id" bigint NOT NULL,
	"project_user_id" bigint NOT NULL,
	CONSTRAINT task_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

/*Общие теги (в данный момент эти теги общие для всех частей(модулей), т.е. у задач, курсов и т.д. одни и теже)*/
CREATE TABLE "tag" (
	"id" serial NOT NULL,
	"name" varchar(25) NOT NULL,
	CONSTRAINT tag_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);


/*Связь задач и тегов*/
CREATE TABLE "tasks_tags_table" (
	"id" serial NOT NULL,
	"task_id" bigint NOT NULL,
	"tag_id" bigint NOT NULL,
	CONSTRAINT tasks_tags_table_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

/*Урок - часть Курса*/
CREATE TABLE "lesson" (
	"id" serial NOT NULL,
	"title" varchar(255) NOT NULL,
	"course_id" bigint NOT NULL,
	"description" TEXT NOT NULL,
	"rating" int NOT NULL,
	CONSTRAINT lesson_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

/*Тестовая часть урока*/
CREATE TABLE "test" (
	"id" serial NOT NULL,
	"title" TEXT NOT NULL,
	"lesson_id" bigint NOT NULL,
	CONSTRAINT test_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);


/*Задание в тесте*/
CREATE TABLE "task_for_test" (
	"id" serial NOT NULL,
	"question" TEXT NOT NULL,
	"test_id" bigint NOT NULL,
	CONSTRAINT task_for_test_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);


/*Вариант ответа в задании теста*/
CREATE TABLE "variant" (
	"id" serial NOT NULL,
	"text" TEXT NOT NULL,
	"correct_answer" bool NOT NULL,
	"task_for_test_id" bigint NOT NULL,
	CONSTRAINT variant_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);


/*Практическая часть урока*/
CREATE TABLE "practise" (
	"id" serial NOT NULL,
	"lesson_id" bigint NOT NULL,
	"description" TEXT NOT NULL,
	CONSTRAINT practise_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);


/*Видео(теоретическая) часть урока*/
CREATE TABLE "video" (
	"id" serial NOT NULL,
	"lesson_id" bigint NOT NULL,
	"link" TEXT NOT NULL,
	CONSTRAINT video_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);


/*Связь пользователя и курса, на который он записан*/
CREATE TABLE "user_course" (
	"id" serial NOT NULL,
	"user_id" bigint NOT NULL,
	"course_id" bigint NOT NULL,
	"date" TIMESTAMP WITH TIME ZONE NOT NULL,
	"progress" int NOT NULL,
	CONSTRAINT user_course_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);


/*Вспомогательная таблица, устанавливающая соответствие между комментарием и его автором*/
CREATE TABLE "user_comments" (
	"id" serial NOT NULL,
	"user_id" bigint NOT NULL,
	"comment_id" bigint NOT NULL, -- в случае цитирования или комментирования комментария
	"date" TIMESTAMP WITH TIME ZONE NOT NULL,
	CONSTRAINT user_comments_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

/*Комментарий пользователя*/
CREATE TABLE "comment" (
  "id" serial NOT NULL,
  "text" TEXT NOT NULL,
  CONSTRAINT comment_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

/*Курс*/
CREATE TABLE "course" (
	"id" serial NOT NULL,
	"title" varchar(255) NOT NULL,
	"description" TEXT NOT NULL,
	"logo" varchar(1000) NOT NULL,
	"rating" real NOT NULL,
	CONSTRAINT course_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);


/*Список рекомендованных курсов для данного курса*/
CREATE TABLE "recomend_courses" (
	"id" serial NOT NULL,
	"course_id" bigint NOT NULL,
	"link" varchar(255) NOT NULL,
	CONSTRAINT recomend_courses_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);


/*Дополнительные данные по курсу*/
CREATE TABLE "statistic" (
	"id" serial NOT NULL,
	"hours" bigint NOT NULL,
	"course_id" bigint NOT NULL,
	CONSTRAINT statistic_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
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

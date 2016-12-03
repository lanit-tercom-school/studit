CREATE TABLE "project" (
	"id" serial NOT NULL,
	"name" varchar(100) NOT NULL,
	"description" TEXT NOT NULL,
	"logo" varchar(1000) NOT NULL,
	CONSTRAINT project_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "course" (
	"id" serial NOT NULL,
	"title" varchar(255) NOT NULL,
	"description" TEXT NOT NULL,
	"logo" varchar(1000) NOT NULL,
	CONSTRAINT course_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "news" (
	"id" serial NOT NULL,
	"title" varchar(255) NOT NULL,
	"text" TEXT NOT NULL,
	CONSTRAINT news_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "author" (
	"id" serial NOT NULL,
	"user_id" bigint NOT NULL,
	CONSTRAINT author_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "project_author" (
	"id" serial NOT NULL,
	"author_id" bigint NOT NULL,
	"project_id" bigint NOT NULL,
	CONSTRAINT project_author_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "lesson" (
	"id" serial NOT NULL,
	"title" varchar(255) NOT NULL,
	"course_id" bigint NOT NULL,
	"description" TEXT NOT NULL,
	CONSTRAINT lesson_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "user" (
	"id" serial NOT NULL,
	"login" varchar(100) NOT NULL,
	"password" bit varying(100) NOT NULL,
	"nickname" varchar(100) NOT NULL,
	"description" TEXT NOT NULL,
	"avatar" varchar(1000) NOT NULL,
	CONSTRAINT user_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "task" (
	"id" serial NOT NULL,
	"title" varchar(100) NOT NULL,
	"description" varchar(255) NOT NULL,
	"project_id" bigint NOT NULL,
	CONSTRAINT task_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "project_user" (
	"id" serial NOT NULL,
	"project_id" bigint NOT NULL,
	"user_id" bigint NOT NULL,
	CONSTRAINT project_user_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "user_contact" (
	"id" serial NOT NULL,
	"contact" varchar(255) NOT NULL,
	"contact_type_id" bigint NOT NULL,
	"user_id" bigint NOT NULL,
	CONSTRAINT user_contact_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "contact_type" (
	"id" serial NOT NULL,
	"type" varchar(100) NOT NULL,
	CONSTRAINT contact_type_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);






ALTER TABLE "author" ADD CONSTRAINT "author_fk0" FOREIGN KEY ("user_id") REFERENCES "user"("id");

ALTER TABLE "project_author" ADD CONSTRAINT "project_author_fk0" FOREIGN KEY ("author_id") REFERENCES "author"("id");
ALTER TABLE "project_author" ADD CONSTRAINT "project_author_fk1" FOREIGN KEY ("project_id") REFERENCES "project"("id");

ALTER TABLE "lesson" ADD CONSTRAINT "lesson_fk0" FOREIGN KEY ("course_id") REFERENCES "course"("id");


ALTER TABLE "task" ADD CONSTRAINT "task_fk0" FOREIGN KEY ("project_id") REFERENCES "project"("id");

ALTER TABLE "project_user" ADD CONSTRAINT "project_user_fk0" FOREIGN KEY ("project_id") REFERENCES "project"("id");
ALTER TABLE "project_user" ADD CONSTRAINT "project_user_fk1" FOREIGN KEY ("user_id") REFERENCES "user"("id");

ALTER TABLE "user_contact" ADD CONSTRAINT "user_contact_fk0" FOREIGN KEY ("contact_type_id") REFERENCES "contact_type"("id");
ALTER TABLE "user_contact" ADD CONSTRAINT "user_contact_fk1" FOREIGN KEY ("user_id") REFERENCES "user"("id");



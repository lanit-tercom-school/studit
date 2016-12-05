CREATE TABLE "user_project" (
	"id" serial NOT NULL,
	"user_id" bigint NOT NULL,
	"project_id" bigint NOT NULL,
	"date" DATE NOT NULL,
	"progress" int NOT NULL,
	CONSTRAINT user_project_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "user_course" (
	"id" serial NOT NULL,
	"user_id" bigint NOT NULL,
	"course_id" bigint NOT NULL,
	"date" DATE NOT NULL,
	"progress" int NOT NULL,
	CONSTRAINT user_course_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "user_comments" (
	"id" serial NOT NULL,
	"user_id" bigint NOT NULL,
	"comment_id" bigint NOT NULL,
	"date" DATE NOT NULL,
	CONSTRAINT user_comments_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);







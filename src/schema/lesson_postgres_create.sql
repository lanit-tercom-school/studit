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



CREATE TABLE "course" (
	"id" serial NOT NULL,
	CONSTRAINT course_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "test" (
	"id" serial NOT NULL,
	"title" TEXT NOT NULL,
	"lesson_id" bigint NOT NULL,
	CONSTRAINT test_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "task_for_test" (
	"id" serial NOT NULL,
	"question" TEXT NOT NULL,
	"test_id" bigint NOT NULL,
	CONSTRAINT task_for_test_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "variant" (
	"id" serial NOT NULL,
	"text" TEXT NOT NULL,
	"correct_answer" bool NOT NULL,
	"task_for_test_id" bigint NOT NULL,
	CONSTRAINT variant_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "practise" (
	"id" serial NOT NULL,
	"lesson_id" bigint NOT NULL,
	"description" TEXT NOT NULL,
	CONSTRAINT practise_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "video" (
	"id" serial NOT NULL,
	"lesson_id" bigint NOT NULL,
	"link" TEXT NOT NULL,
	CONSTRAINT video_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



ALTER TABLE "lesson" ADD CONSTRAINT "lesson_fk0" FOREIGN KEY ("course_id") REFERENCES "course"("id");


ALTER TABLE "test" ADD CONSTRAINT "test_fk0" FOREIGN KEY ("lesson_id") REFERENCES "lesson"("id");

ALTER TABLE "task_for_test" ADD CONSTRAINT "task_for_test_fk0" FOREIGN KEY ("test_id") REFERENCES "test"("id");

ALTER TABLE "variant" ADD CONSTRAINT "variant_fk0" FOREIGN KEY ("task_for_test_id") REFERENCES "task_for_test"("id");

ALTER TABLE "practise" ADD CONSTRAINT "practise_fk0" FOREIGN KEY ("lesson_id") REFERENCES "lesson"("id");

ALTER TABLE "video" ADD CONSTRAINT "video_fk0" FOREIGN KEY ("lesson_id") REFERENCES "lesson"("id");


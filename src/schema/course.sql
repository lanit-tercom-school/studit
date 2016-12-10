CREATE TABLE "course" (
	"id" serial NOT NULL,
	"title" varchar(255) NOT NULL,
	"discription" TEXT NOT NULL,
	"rating" real NOT NULL,
	CONSTRAINT course_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "recomend_courses" (
	"id" serial NOT NULL,
	"course_id" bigint NOT NULL,
	"link" serial NOT NULL,
	CONSTRAINT recomend_courses_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "statistic" (
	"id" serial NOT NULL,
	"hours" serial NOT NULL,
	"course_id" serial NOT NULL,
	CONSTRAINT statistic_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);




ALTER TABLE "recomend_courses" ADD CONSTRAINT "recomend_courses_fk0" FOREIGN KEY ("course_id") REFERENCES "course"("id");

ALTER TABLE "statistic" ADD CONSTRAINT "statistic_fk0" FOREIGN KEY ("course_id") REFERENCES "course"("id");


CREATE TABLE "course" (
	"id" serial,
	"title" varchar(255),
	"discription" TEXT,
	"rating" double,
	CONSTRAINT course_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "recomend_courses" (
	"id" serial,
	"course_id" bigint,
	"link" serial,
	CONSTRAINT recomend_courses_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "statistic" (
	"id" serial,
	"hours" serial,
	"course_id" serial,
	CONSTRAINT statistic_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);




ALTER TABLE "recomend_courses" ADD CONSTRAINT "recomend_courses_fk0" FOREIGN KEY ("course_id") REFERENCES "course"("id");

ALTER TABLE "statistic" ADD CONSTRAINT "statistic_fk0" FOREIGN KEY ("course_id") REFERENCES "course"("id");


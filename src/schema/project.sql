CREATE TABLE "project" (
	"id" serial NOT NULL,
	"description" TEXT NOT NULL,
	"tasks" TEXT NOT NULL,
	"marerials" TEXT NOT NULL,
	"news" TEXT NOT NULL,
	CONSTRAINT project_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);





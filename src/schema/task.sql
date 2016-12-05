CREATE TABLE "task" (
	"id" serial NOT NULL,
	"title" varchar(100) NOT NULL,
	"description" varchar(255) NOT NULL,
	"numberOfTask" int NOT NULL,
	"tags" varchar NOT NULL,
	"priority" int NOT NULL,
	"project_id" bigint NOT NULL,
	"project_author_id" bigint NOT NULL,
	"project_user_id" bigint NOT NULL,
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



CREATE TABLE "project_author" (
	"id" serial NOT NULL,
	"author_id" bigint NOT NULL,
	"project_id" bigint NOT NULL,
	CONSTRAINT project_author_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "project" (
	"id" serial NOT NULL,
	"name" varchar(100) NOT NULL,
	"description" varchar(100) NOT NULL,
	CONSTRAINT project_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "tag" (
	"id" serial NOT NULL,
	"name" varchar(25) NOT NULL,
	CONSTRAINT tag_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "tags_table" (
	"id" serial NOT NULL,
	"task_id" bigint NOT NULL,
	"tag_id" bigint NOT NULL,
	CONSTRAINT tags_table_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



ALTER TABLE "task" ADD CONSTRAINT "task_fk0" FOREIGN KEY ("project_id") REFERENCES "project"("id");
ALTER TABLE "task" ADD CONSTRAINT "task_fk1" FOREIGN KEY ("project_author_id") REFERENCES "project_author"("id");
ALTER TABLE "task" ADD CONSTRAINT "task_fk2" FOREIGN KEY ("project_user_id") REFERENCES "project_user"("id");

ALTER TABLE "project_user" ADD CONSTRAINT "project_user_fk0" FOREIGN KEY ("project_id") REFERENCES "project"("id");

ALTER TABLE "project_author" ADD CONSTRAINT "project_author_fk0" FOREIGN KEY ("project_id") REFERENCES "project"("id");



ALTER TABLE "tags_table" ADD CONSTRAINT "tags_table_fk0" FOREIGN KEY ("task_id") REFERENCES "task"("id");
ALTER TABLE "tags_table" ADD CONSTRAINT "tags_table_fk1" FOREIGN KEY ("tag_id") REFERENCES "tag"("id");


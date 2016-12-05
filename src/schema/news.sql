CREATE TABLE "news" (
	"id" serial NOT NULL,
	"title" varchar(255) NOT NULL,
	"description" TEXT NOT NULL,
	"date" DATE NOT NULL,
	CONSTRAINT news_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "news_tags" (
	"id" serial NOT NULL,
	"text" TEXT NOT NULL,
	CONSTRAINT news_tags_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "news_news_tags" (
	"id" serial NOT NULL,
	"news_id" bigint NOT NULL,
	"news_tags_id" bigint NOT NULL,
	CONSTRAINT news_news_tags_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);





ALTER TABLE "news_news_tags" ADD CONSTRAINT "news_news_tags_fk0" FOREIGN KEY ("news_id") REFERENCES "news"("id");
ALTER TABLE "news_news_tags" ADD CONSTRAINT "news_news_tags_fk1" FOREIGN KEY ("news_tags_id") REFERENCES "news_tags"("id");


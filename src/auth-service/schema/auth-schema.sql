CREATE TABLE "user" (
  "id"               SERIAL       NOT NULL,
  "login"            VARCHAR(100) NOT NULL,
  "password"         VARCHAR(100) NOT NULL,
  "permission_level" INT          NOT NULL,
  CONSTRAINT user_pk PRIMARY KEY ("id")
) WITH (
OIDS = FALSE
);

CREATE TABLE "wait_for_activation" (
  "id"       SERIAL NOT NULL,
  "user_id"  INT    NOT NULL,
  "nickname" VARCHAR(100),
  CONSTRAINT wfa PRIMARY KEY ("id")
) WITH (
OIDS = FALSE
);
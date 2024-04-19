CREATE TABLE "users" (
  "id" bigserial  PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "todos" (
  "id" bigserial  PRIMARY KEY,
  "title" varchar NOT NULL,
  "user_id" bigint NOT NULL,
  "done" boolean DEFAULT (false),
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "todos" ("user_id");

ALTER TABLE "todos" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

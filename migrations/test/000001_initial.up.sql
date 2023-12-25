CREATE TABLE "users" (
                         "id" integer PRIMARY KEY autoincrement,
                         "email" varchar NOT NULL ,
                         "firstname" varchar,
                         "lastname" varchar
);

CREATE TABLE "posts" (
                         "id" bigserial PRIMARY KEY,
                         "content" text
);

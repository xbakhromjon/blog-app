CREATE TABLE "users" (
                         "id" bigserial PRIMARY KEY,
                         "email" varchar NOT NULL ,
                         "firstname" varchar,
                         "lastname" varchar
);

CREATE TABLE "posts" (
                         "id" bigserial PRIMARY KEY,
                         "content" text
);

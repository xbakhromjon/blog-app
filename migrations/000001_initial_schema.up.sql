CREATE TABLE "users" (
                         "id" bigserial PRIMARY KEY,
                         "username" varchar UNIQUE NOT NULL,
                         "password" varchar NOT NULL
);

CREATE TABLE "posts" (
                         "id" bigserial PRIMARY KEY,
                         "content" text
);

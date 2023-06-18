CREATE TABLE "authors" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "date_of_birth" DATE NOT NULL
);
CREATE TABLE "books" (
    "id" SERIAL PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL,
    "author_id" INT NOT NULL,
    "publish_date" DATE NOT NULL,
    FOREIGN KEY ("author_id") REFERENCES "authors"("id")
);
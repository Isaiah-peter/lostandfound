CREATE TYPE "post_type_status" AS ENUM (
  'lost',
  'found'
);

CREATE TYPE "item_status" AS ENUM (
  'claimed',
  'unclamed'
);

CREATE TABLE "users" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "address" varchar NOT NULL,
  "contact" varchar NOT NULL,
  "username" varchar NOT NULL,
  "user_image" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "title" varchar NOT NULL,
  "discription" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "lost_items" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "category_id" int,
  "founder_id" int NOT NULL,
  "title" varchar NOT NULL,
  "discription" varchar NOT NULL,
  "date" timestamp NOT NULL,
  "time" varchar NOT NULL,
  "location" varchar NOT NULL NOT NULL,
  "post_type" post_type_status NOT NULL,
  "status" item_status NOT NULL,
  "remark" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "lost_items_images" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "lost_item_id" int NOT NULL,
  "lost_item_image" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE INDEX ON "lost_items" ("category_id");

CREATE INDEX ON "lost_items" ("founder_id");

CREATE INDEX ON "lost_items_images" ("lost_item_id");

ALTER TABLE "lost_items" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "lost_items" ADD FOREIGN KEY ("founder_id") REFERENCES "users" ("id");

ALTER TABLE "lost_items_images" ADD FOREIGN KEY ("lost_item_id") REFERENCES "lost_items" ("id");
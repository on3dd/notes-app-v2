CREATE TABLE "notes" (
	"id" serial NOT NULL,
	"author_id" int NOT NULL,
	"category_id" int NOT NULL,
	"subject_id" int NOT NULL,
	"teacher_id" int NOT NULL,
	"posted_at" TIMESTAMP NOT NULL,
	"title" TEXT NOT NULL,
	"descirption" TEXT NOT NULL,
	"link" TEXT NOT NULL UNIQUE,
	CONSTRAINT "notes_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "users" (
	"id" serial NOT NULL,
	"name" TEXT NOT NULL UNIQUE,
	"password" TEXT,
	"email" varchar(255) NOT NULL,
	"about" TEXT NOT NULL,
	CONSTRAINT "users_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "categories" (
	"id" serial NOT NULL,
	"name" TEXT NOT NULL,
	"description" TEXT,
	CONSTRAINT "categories_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "subjects" (
	"id" serial NOT NULL,
	"category_id" int NOT NULL,
	"name" TEXT NOT NULL,
	CONSTRAINT "subjects_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "teachers" (
	"id" serial NOT NULL,
	"category_id" int NOT NULL,
	"name" TEXT NOT NULL UNIQUE,
	CONSTRAINT "teachers_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



ALTER TABLE "notes" ADD CONSTRAINT "notes_fk0" FOREIGN KEY ("author_id") REFERENCES "users"("id");
ALTER TABLE "notes" ADD CONSTRAINT "notes_fk1" FOREIGN KEY ("category_id") REFERENCES "categories"("id");
ALTER TABLE "notes" ADD CONSTRAINT "notes_fk2" FOREIGN KEY ("subject_id") REFERENCES "subjects"("id");
ALTER TABLE "notes" ADD CONSTRAINT "notes_fk3" FOREIGN KEY ("teacher_id") REFERENCES "teachers"("id");



ALTER TABLE "subjects" ADD CONSTRAINT "subjects_fk0" FOREIGN KEY ("category_id") REFERENCES "categories"("id");

ALTER TABLE "teachers" ADD CONSTRAINT "teachers_fk0" FOREIGN KEY ("category_id") REFERENCES "categories"("id");

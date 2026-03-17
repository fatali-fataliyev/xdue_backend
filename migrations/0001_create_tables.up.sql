CREATE TABLE
  IF NOT EXISTS "privacy_policy" (
    "content" text,
    "updated_at" timestamp DEFAULT (now())
  );


CREATE TABLE
  IF NOT EXISTS "dev" (
    "id" uuid PRIMARY KEY,
    "email" varchar(255) NOT NULL,
    "password_hash" varchar(255)
  );


CREATE TABLE
  IF NOT EXISTS "feedbacks" (
    "id" uuid PRIMARY KEY,
    "category" varchar(255) NOT NULL,
    "app_rating" int NOT NULL,
    "feedback" text NOT NULL,
    "name" varchar(255),
    "email" varchar(255)
  );


CREATE TABLE
  IF NOT EXISTS "dev_session" (
    "id" uuid PRIMARY KEY,
    "token" varchar(255) NOT NULL,
    "dev_id" uuid NOT NULL,
    "expire_at" timestamp
  );


CREATE TABLE
  IF NOT EXISTS "users" (
    "id" uuid PRIMARY KEY,
    "name" varchar(255) NOT NULL,
    "email" varchar(255) NOT NULL,
    "password_hash" varchar(255) NOT NULL,
    "created_at" timestamp DEFAULT (now())
  );


CREATE TABLE
  IF NOT EXISTS "groups" (
    "id" uuid PRIMARY KEY,
    "name" varchar(255) NOT NULL,
    "type" varchar(255) NOT NULL,
    "currency_iso" char(3) NOT NULL,
    "created_by" uuid NOT NULL,
    "updated_at" timestamp DEFAULT (now()),
    "created_at" timestamp DEFAULT (now())
  );


CREATE TABLE
  IF NOT EXISTS "pending_members" (
    "id" uuid PRIMARY KEY,
    "group_id" uuid NOT NULL,
    "sender_id" uuid NOT NULL,
    "sent_at" timestamp DEFAULT (now())
  );


CREATE TABLE
  IF NOT EXISTS "group_members" (
    "id" uuid PRIMARY KEY,
    "group_id" uuid NOT NULL,
    "user_id" uuid NOT NULL,
    "joined_at" timestamp DEFAULT (now())
  );


CREATE TABLE
  IF NOT EXISTS "expenses" (
    "id" uuid PRIMARY KEY,
    "title" varchar(255) DEFAULT '',
    "total_amount" bigint NOT NULL,
    "currency_iso" char(3) NOT NULL,
    "group_id" uuid NOT NULL,
    "created_by" uuid NOT NULL,
    "created_at" timestamp DEFAULT (now()),
    "updated_at" timestamp DEFAULT (now()),
    "note" text DEFAULT '',
    "is_deleted" boolean DEFAULT false
  );


CREATE TABLE
  IF NOT EXISTS "expense_payments" (
    "id" uuid PRIMARY KEY,
    "expense_id" uuid NOT NULL,
    "paid_amount" bigint NOT NULL,
    "payer_id" uuid NOT NULL
  );


CREATE TABLE
  IF NOT EXISTS "expense_splits" (
    "id" uuid PRIMARY KEY,
    "group_id" uuid NOT NULL,
    "user_id" uuid NOT NULL,
    "expense_id" uuid NOT NULL,
    "split_method" varchar(255) NOT NULL,
    "method_value" bigint NOT NULL,
    "is_exclude" bool DEFAULT false
  );


CREATE TABLE
  IF NOT EXISTS "settle_ups"(
    "id" uuid PRIMARY KEY NOT NULL,
    "amount" bigint NOT NULL,
    "expense_id" uuid NOT NULL,
    "payer_id" uuid NOT NULL,
    "receiver_id" uuid NOT NULL,
    "note" text DEFAULT '',
    "created_by" uuid NOT NULL,
    "created_at" timestamp DEFAULT (now())
  );


CREATE TABLE
  IF NOT EXISTS "sessions" (
    "id" uuid PRIMARY KEY,
    "token" varchar(255) NOT NULL,
    "user_id" uuid NOT NULL,
    "expire_at" timestamp NOT NULL
  );


CREATE TABLE
  IF NOT EXISTS "notifications" (
    "id" uuid PRIMARY KEY,
    "content" varchar(255),
    "receiver_id" uuid NOT NULL,
    "created_at" timestamp DEFAULT (now()),
    "is_read" bool DEFAULT false
  );


ALTER TABLE
  "groups"
ADD
  FOREIGN KEY ("created_by") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "expenses"
ADD
  FOREIGN KEY ("group_id") REFERENCES "groups" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "group_members"
ADD
  FOREIGN KEY ("group_id") REFERENCES "groups" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "group_members"
ADD
  FOREIGN KEY ("user_id") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "sessions"
ADD
  FOREIGN KEY ("user_id") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "notifications"
ADD
  FOREIGN KEY ("receiver_id") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "expenses"
ADD
  FOREIGN KEY ("created_by") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "settle_ups"
ADD
  FOREIGN KEY ("expense_id") REFERENCES "expenses" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "settle_ups"
ADD
  FOREIGN KEY ("payer_id") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "settle_ups"
ADD
  FOREIGN KEY ("created_by") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "settle_ups"
ADD
  FOREIGN KEY ("receiver_id") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "expense_payments"
ADD
  FOREIGN KEY ("expense_id") REFERENCES "expenses" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "expense_payments"
ADD
  FOREIGN KEY ("payer_id") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "expense_splits"
ADD
  FOREIGN KEY ("group_id") REFERENCES "groups" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "expense_splits"
ADD
  FOREIGN KEY ("user_id") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "expense_splits"
ADD
  FOREIGN KEY ("expense_id") REFERENCES "expenses" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "dev_session"
ADD
  FOREIGN KEY ("dev_id") REFERENCES "dev" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "pending_members"
ADD
  FOREIGN KEY ("group_id") REFERENCES "groups" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "pending_members"
ADD
  FOREIGN KEY ("sender_id") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;
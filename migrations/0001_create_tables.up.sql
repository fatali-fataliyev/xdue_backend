--developer section tables
CREATE TABLE
  "privacy_policy" (
    "content" text,
    "updated_at" timestamp DEFAULT (now())
  );


CREATE TABLE
  "dev" (
    "id" uuid PRIMARY KEY,
    "email" varchar(255) NOT NULL,
    "password_hash" varchar(255)
  );


CREATE TABLE
  "feedbacks" (
    "id" uuid PRIMARY KEY,
    "category" varchar(255) NOT NULL,
    "app_rating" int NOT NULL,
    "feedback" text NOT NULL,
    "name" varchar(255),
    "email" varchar(255)
  );


CREATE TABLE
  "dev_session" (
    "id" uuid PRIMARY KEY,
    "token" varchar(255) NOT NULL,
    "dev_id" uuid NOT NULL,
    "expire_at" TIMESTAMPTZ
  );


--user section tables
CREATE TABLE
  "groups" (
    "id" uuid PRIMARY KEY,
    "name" varchar(255) NOT NULL,
    "type" varchar(255) NOT NULL,
    "currency_iso" char(3) NOT NULL,
    "updated_at" timestamp DEFAULT (now()),
    "created_at" timestamp DEFAULT (now())
  );


CREATE TABLE
  "pending_members" (
    "id" uuid PRIMARY KEY,
    "group_id" uuid NOT NULL,
    "name" varchar(255) NOT NULL,
    "sent_at" timestamp DEFAULT (now())
  );


CREATE TABLE
  "members" (
    "id" uuid PRIMARY KEY,
    "name" varchar(255) NOT NULL,
    "backup_key" varchar(255) NOT NULL,
    "group_id" uuid NOT NULL,
    "role" varchar(255) NOT NULL
  );


CREATE TABLE
  "expenses" (
    "id" uuid PRIMARY KEY,
    "title" varchar(255) DEFAULT '',
    "total_amount" bigint NOT NULL,
    "currency_iso" char(3) NOT NULL,
    "group_id" uuid NOT NULL,
    "created_by" uuid NOT NULL,
    "created_at" timestamp DEFAULT (now()),
    "updated_at" timestamp DEFAULT (now())
  );


CREATE TABLE
  "deleted_expenses" (
    "id" uuid PRIMARY KEY,
    "title" varchar(255) DEFAULT '',
    "total_amount" bigint NOT NULL,
    "currency_iso" char(3) NOT NULL,
    "group_id" uuid NOT NULL,
    "deleted_by" uuid NOT NULL,
    "created_at" timestamp DEFAULT (now()),
    "deleted_at" timestamp NOT NULL
  );


CREATE TABLE
  "expense_payments" (
    "id" uuid PRIMARY KEY,
    "expense_id" uuid NOT NULL,
    "paid_amount" bigint NOT NULL,
    "payer_id" uuid NOT NULL
  );


CREATE TABLE
  "expense_splits" (
    "id" uuid PRIMARY KEY,
    "group_id" uuid NOT NULL,
    "member_id" uuid NOT NULL,
    "expense_id" uuid NOT NULL,
    "split_method" varchar(255) NOT NULL,
    "method_value" bigint NOT NULL,
    "is_exclude" bool DEFAULT false
  );


CREATE TABLE
  "sessions" (
    "id" uuid PRIMARY KEY,
    "token" varchar(255) NOT NULL,
    "member_id" uuid NOT NULL,
    "expire_at" TIMESTAMPTZ NOT NULL
  );


CREATE TABLE
  "notifications" (
    "id" uuid PRIMARY KEY,
    "content" varchar(255),
    "receiver_id" uuid NOT NULL,
    "created_at" timestamp DEFAULT (now()),
    "is_read" bool DEFAULT false
  );


ALTER TABLE
  "members"
ADD
  FOREIGN KEY ("group_id") REFERENCES "groups" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "expenses"
ADD
  FOREIGN KEY ("group_id") REFERENCES "groups" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "expenses"
ADD
  FOREIGN KEY ("created_by") REFERENCES "members" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "notifications"
ADD
  FOREIGN KEY ("receiver_id") REFERENCES "members" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "sessions"
ADD
  FOREIGN KEY ("member_id") REFERENCES "members" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "pending_members"
ADD
  FOREIGN KEY ("group_id") REFERENCES "groups" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "expense_payments"
ADD
  FOREIGN KEY ("expense_id") REFERENCES "expenses" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "expense_payments"
ADD
  FOREIGN KEY ("payer_id") REFERENCES "members" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "expense_splits"
ADD
  FOREIGN KEY ("group_id") REFERENCES "groups" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "expense_splits"
ADD
  FOREIGN KEY ("member_id") REFERENCES "members" ("id") DEFERRABLE INITIALLY IMMEDIATE;


ALTER TABLE
  "expense_splits"
ADD
  FOREIGN KEY ("expense_id") REFERENCES "expenses" ("id") DEFERRABLE INITIALLY IMMEDIATE;
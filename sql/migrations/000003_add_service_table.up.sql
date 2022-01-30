CREATE TABLE IF NOT EXISTS services(
                                     "id" uuid PRIMARY KEY,
                                     "name" varchar(64) not null unique,
                                     "description" text not null,
                                     "status" varchar(64) not null,
                                     "status_since" timestamp with time zone DEFAULT now(),
                                     "created_at" timestamp with time zone DEFAULT now(),
                                     "updated_at" timestamp with time zone DEFAULT now(),
                                     "pages" text[] NOT NULL DEFAULT '{}'::text[],
                                     "group_id" varchar(64)
);
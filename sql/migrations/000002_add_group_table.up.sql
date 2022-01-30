CREATE TABLE IF NOT EXISTS groups(
                                     "id" uuid PRIMARY KEY,
                                     "name" varchar(64) not null unique,
                                     "description" text not null,
                                     "expanded" boolean,
                                     "created_at" timestamp with time zone DEFAULT now(),
                                     "updated_at" timestamp with time zone DEFAULT now()
);
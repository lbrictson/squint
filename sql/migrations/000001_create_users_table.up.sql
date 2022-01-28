CREATE TABLE IF NOT EXISTS users (
                                     "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
                                     "username" varchar(64) not null unique,
                                     "hashed_password" text not null,
                                     "role" varchar(64) not null,
                                     "api_key" varchar(64) not null,
                                     "created_at" timestamp with time zone DEFAULT now(),
                                     "updated_at" timestamp with time zone DEFAULT now()
);

CREATE UNIQUE INDEX users_username ON users (username);
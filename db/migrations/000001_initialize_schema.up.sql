BEGIN;

CREATE TABLE IF NOT EXISTS comments(
   id uuid PRIMARY KEY,
   book VARCHAR (50) NOT NULL,
   comment VARCHAR (500) NOT NULL,
   ip_address VARCHAR (50) NOT NULL,
   created_at timestamp WITH time zone NOT NULL
);

COMMIT;
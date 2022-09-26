BEGIN;

CREATE TABLE IF NOT EXISTS book_comments(
   id uuid PRIMARY KEY,
   book_id VARCHAR (50) NOT NULL,
   ip_address VARCHAR (50) NOT NULL,
   created timestamp WITH time zone NOT NULL,
   comment VARCHAR (500) NOT NULL
);

COMMIT;
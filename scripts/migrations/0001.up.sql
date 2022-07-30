BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS customers (
  id uuid NOT NULL DEFAULT uuid_generate_v4(),
  first_name character varying(50) NOT NULL,
  last_name character varying(100) NOT NULL,
  birth_date date NOT NULL,
  deleted_at timestamp NULL,
  updated_at timestamp NOT NULL DEFAULT now(),
  created_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS migrations
(
    name character varying(100) NOT NULL PRIMARY KEY,
    applied_at timestamp without time zone NOT NULL
);

INSERT INTO migrations VALUES ('0001.up.sql', NOW());
COMMIT;
END;
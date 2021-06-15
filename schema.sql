-- The PostgreSQL schema
CREATE DATABASE ocp_presentation_api;

-- Make sure we are using the `ocp_presentation_api` database
\c ocp_presentation_api;

CREATE TABLE presentation (
  id          SERIAL PRIMARY KEY,
  lesson_id   INT,
  user_id     INT,
  name        VARCHAR,
  description VARCHAR
);

CREATE TYPE content_type AS ENUM (
  'Document',
  'Video',
  'Question',
  'Task'
);

CREATE TABLE slide (
  id              SERIAL PRIMARY KEY,
  presentation_id INT,
  number          INT,
  type            content_type
);

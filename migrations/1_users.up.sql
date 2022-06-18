CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  uname VARCHAR UNIQUE NOT NULL,
  email VARCHAR UNIQUE NOT NULL,
  password VARCHAR NOT NULL,
  phone VARCHAR,
  profile_picture VARCHAR,
  full_name VARCHAR,
  first_name VARCHAR,
  last_name VARCHAR,
  nickname VARCHAR,
  user_address VARCHAR,
  created_at VARCHAR DEFAULT CURRENT_TIMESTAMP,
  updated_at VARCHAR DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS contact (
  id SERIAL PRIMARY KEY,
  uname VARCHAR NOT NULL,
  email VARCHAR ,
  phone VARCHAR NOT NULL,
  messages VARCHAR,
  relationship VARCHAR,
  priorities INT , 
  profile_picture VARCHAR,
  alternative_name  VARCHAR,
  created_at VARCHAR DEFAULT CURRENT_TIMESTAMP,
  updated_at VARCHAR DEFAULT CURRENT_TIMESTAMP,
  user_id BIGINT REFERENCES users(id)
);


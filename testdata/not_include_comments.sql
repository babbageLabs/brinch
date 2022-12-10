DROP TABLE IF EXISTS articles;
CREATE TABLE articles (
  id BIGINT,
  user_id BIGINT,
  title VARCHAR(255),
  content TEXT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);
INSERT INTO articles (
  id, user_id, title, content, created_at, updated_at
) VALUES (
  1, 1, 'Self Introduction', '-- sqlfile --\nI''m sqlfile.', now(), now()
);
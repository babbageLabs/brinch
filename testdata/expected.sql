DROP TABLE IF EXISTS articles;
CREATE TABLE articles
(
    id         BIGINT,
    user_id    BIGINT,
    title      VARCHAR(255),
    content    TEXT
);

INSERT INTO articles (id, user_id, title, content)
VALUES (1, 1, 'Self Introduction', '-- sqlfile --\nI''m sqlfile.');
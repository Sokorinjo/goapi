DROP TABLE IF EXISTS users;
CREATE TABLE users (
  user_id INT AUTO_INCREMENT PRIMARY KEY,
  user_name VARCHAR(30) NOT NULL,
  user_pass VARCHAR(30) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users 
  (user_name, user_pass)
VALUES
  ("vlada", "darkcidator"),
  ("Kuroppi", "1234"),
  ("vlasta peder", "peder");

DROP TABLE IF EXISTS posts;
CREATE TABLE posts(
  post_id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  user_id INT,
  content TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO posts
  (user_id, content)
VALUES 
  ('1', 'jak gas'),
  ('1', 'Najjaci sam.'),
  ('2', 'imam najlepseg decka'),
  ('3', 'ja sam pedercina <3');
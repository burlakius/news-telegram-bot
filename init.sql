CREATE DATABASE IF NOT EXISTS tgbot_db;
USE tgbot_db;

-- News types table
CREATE TABLE IF NOT EXISTS news_types (
  news_type_id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

-- News table
CREATE TABLE IF NOT EXISTS news (
  news_id INT AUTO_INCREMENT PRIMARY KEY,
  news_type INT NOT NULL,
  news_chat_id BIGINT NOT NULL,
  news_message_id INT NOT NULL,
  publication_date DATETIME NOT NULL,
  CONSTRAINT fk_news_types
    FOREIGN KEY (news_type)
    REFERENCES news_types(news_type_id)
    ON DELETE CASCADE
);

-- Users table
CREATE TABLE IF NOT EXISTS users (
  user_id INT AUTO_INCREMENT PRIMARY KEY,
  news_type_id INT DEFAULT 1,
  CONSTRAINT fk_users_news_types
    FOREIGN KEY (news_type_id)
    REFERENCES news_types(news_type_id)
    ON DELETE CASCADE
);

-- User questions table
CREATE TABLE IF NOT EXISTS user_questions (
  question_id INT AUTO_INCREMENT PRIMARY KEY,
  question_chat_id BIGINT NOT NULL,
  question_message_id INT NOT NULL
);

-- Message table of answers to the question
CREATE TABLE IF NOT EXISTS admin_answers (
  answer_id INT AUTO_INCREMENT PRIMARY KEY,
  answer_chat_id BIGINT NOT NULL,
  answer_message_id INT NOT NULL
);

-- Admin chats table
CREATE TABLE IF NOT EXISTS chat_admins (
  admin_chat_id INT AUTO_INCREMENT PRIMARY KEY,
  chat_id BIGINT NOT NULL,
  UNIQUE (chat_id)
);

INSERT INTO news_types (name) VALUES ('regular'), ('important');

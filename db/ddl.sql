-- TODO: とりあえずuserテーブルだけ書いておいたよー。ここに追記する形でDDL定義してくれたら嬉しい！
CREATE TABLE users (
    id   int          auto_increment PRIMARY KEY ,
    uid  VARCHAR(256) NOT NULL,
    name VARCHAR(100) NOT NULL
);
CREATE INDEX idx_users ON users(uid);
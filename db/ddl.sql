CREATE TABLE users (
    id      INT            AUTO_INCREMENT PRIMARY KEY,
    uid     VARCHAR(256)   NOT NULL,
    name    VARCHAR(100)   NOT NULL
);
CREATE INDEX idx_users ON users(uid);

CREATE TABLE challenge_themes (
    id              INT             AUTO_INCREMENT PRIMARY KEY,
    title           VARCHAR(100)    NOT NULL,
    description     VARCHAR(1024)   NOT NULL,
    user_id         INT             NOT NULL,
    recordable      TINYINT(1)      NOT NULL,
    is_int          TINYINT(1),
    unit            VARCHAR(64),
    ranking_type    ENUM('ASC', 'DESC', 'FAV'),
    created_at      DATETIME        NOT NULL
);
CREATE INDEX idx_challenge_themes ON challenge_themes(created_at);

CREATE TABLE challenge_records (
    id                  INT             AUTO_INCREMENT PRIMARY KEY,
    content             VARCHAR(512)    NOT NULL,
    comment             VARCHAR(1024)   NOT NULL,
    challenge_theme_id  INT             NOT NULL,
    user_id             INT             NOT NULL,
    record              FLOAT,
    created_at          DATETIME        NOT NULL
);
CREATE INDEX idx_challenge_records_01 ON challenge_records(challenge_theme_id);
CREATE INDEX idx_challenge_records_02 ON challenge_records(created_at);

CREATE TABLE favorites (
    id                  INT     AUTO_INCREMENT PRIMARY KEY,
    user_id             INT     NOT NULL,
    challenge_record_id INT     NOT NULL
);

CREATE TABLE tags (
    id      INT             AUTO_INCREMENT PRIMARY KEY,
    text    VARCHAR(128)    NOT NULL UNIQUE
);

CREATE TABLE theme_tag_relation (
    id                  INT     AUTO_INCREMENT PRIMARY KEY,
    challenge_theme_id  INT     NOT NULL,
    tag_id              INT     NOT NULL
);

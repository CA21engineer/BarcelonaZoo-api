CREATE TABLE users
(
    id   INT AUTO_INCREMENT PRIMARY KEY,
    uid  VARCHAR(256) NOT NULL,
    name VARCHAR(100) NOT NULL,
    icon VARCHAR(256)
);
CREATE INDEX idx_users ON users (uid);

CREATE TABLE challenge_themes
(
    id           INT AUTO_INCREMENT PRIMARY KEY,
    title        TEXT       NOT NULL,
    content      TEXT       NOT NULL,
    description  TEXT       NOT NULL,
    user_id      INT        NOT NULL,
    recordable   TINYINT(1) NOT NULL,
    is_int       TINYINT(1),
    unit         VARCHAR(64),
    ranking_type TINYINT(3),
    created_at   DATETIME   NOT NULL,
    CONSTRAINT fk_challenge_theme_user_id
        FOREIGN KEY (user_id)
            REFERENCES users (id)
            ON DELETE RESTRICT ON UPDATE RESTRICT
);
CREATE INDEX idx_challenge_themes ON challenge_themes (created_at);

CREATE TABLE challenge_records
(
    id                 INT AUTO_INCREMENT PRIMARY KEY,
    content            TEXT     NOT NULL,
    comment            TEXT     NOT NULL,
    challenge_theme_id INT      NOT NULL,
    user_id            INT      NOT NULL,
    record             FLOAT,
    created_at         DATETIME NOT NULL,
    CONSTRAINT fk_challenge_record_challenge_theme_id
        FOREIGN KEY (challenge_theme_id)
            REFERENCES challenge_themes (id)
            ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT fk_challenge_record_user_id
        FOREIGN KEY (user_id)
            REFERENCES users (id)
            ON DELETE RESTRICT ON UPDATE RESTRICT
);
CREATE INDEX idx_challenge_records ON challenge_records (challenge_theme_id, created_at);

CREATE TABLE favorites
(
    id                  INT AUTO_INCREMENT PRIMARY KEY,
    user_id             INT NOT NULL,
    challenge_record_id INT NOT NULL,
    CONSTRAINT fk_favorite_user_id
        FOREIGN KEY (user_id)
            REFERENCES users (id)
            ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT fk_favorite_challenge_record_id
        FOREIGN KEY (challenge_record_id)
            REFERENCES challenge_records (id)
            ON DELETE RESTRICT ON UPDATE RESTRICT
);

CREATE TABLE tags
(
    id   INT AUTO_INCREMENT PRIMARY KEY,
    text VARCHAR(128) NOT NULL UNIQUE
);

CREATE TABLE theme_tag_relations
(
    id                 INT AUTO_INCREMENT PRIMARY KEY,
    challenge_theme_id INT NOT NULL,
    tag_id             INT NOT NULL,
    CONSTRAINT fk_theme_tag_relation_challenge_theme_id
        FOREIGN KEY (challenge_theme_id)
            REFERENCES challenge_themes (id)
            ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT fk_theme_tag_relation_tag_id
        FOREIGN KEY (tag_id)
            REFERENCES tags (id)
            ON DELETE RESTRICT ON UPDATE RESTRICT
);

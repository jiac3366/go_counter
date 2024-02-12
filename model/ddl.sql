-- 用户信息表
CREATE TABLE user
(
    user_id  bigint NOT NULL AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL DEFAULT '',
    PRIMARY KEY (user_id)
);

-- 文章信息表
CREATE TABLE article
(
    article_id bigint NOT NULL AUTO_INCREMENT,
    title      VARCHAR(255) NOT NULL DEFAULT '',
    content    TEXT         NOT NULL DEFAULT '',
    author_id  bigint       NOT NULL,
    PRIMARY KEY (article_id)
);

-- 点赞记录表
CREATE TABLE likes_record
(
    like_id    bigint NOT NULL AUTO_INCREMENT,
    user_id    bigint NOT NULL,
    article_id bigint NOT NULL,
    created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    PRIMARY KEY (like_id)
    UNIQUE KEY idx_uniq_user_id (user_id)
);


-- 收藏记录表
CREATE TABLE favorites_record
(
    favorite_id bigint NOT NULL AUTO_INCREMENT,
    user_id     bigint NOT NULL,
    article_id  bigint NOT NULL,
    created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    PRIMARY KEY (favorite_id)
    UNIQUE KEY idx_uniq_user_id (user_id)
);

-- 计数表
CREATE TABLE count_meta
(
    id  bigint NOT NULL,
    count bigint NOT NULL,
    type VARCHAR(32) NOT NULL,
    PRIMARY KEY (`id`, `type`)
);

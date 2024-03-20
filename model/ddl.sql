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
    UNIQUE KEY idx_uniq_user_id_article_id (user_id, article_id),
    PRIMARY KEY (like_id)
);


-- 收藏记录表
CREATE TABLE favorites_record
(
    favorite_id bigint NOT NULL AUTO_INCREMENT,
    user_id     bigint NOT NULL,
    article_id  bigint NOT NULL,
    created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    UNIQUE KEY idx_uniq_user_id_article_id (user_id, article_id),
    PRIMARY KEY (favorite_id)
);

-- 计数表
CREATE TABLE count_meta
(
    id  bigint NOT NULL,
    business_id bigint NOT NULL,  -- same as article_id
    count bigint NOT NULL,
    types VARCHAR(32) NOT NULL,
    UNIQUE KEY idx_uniq_business_id (business_id, types), -- article_id_1, "FAVORITE" 代表 id 为 1 的文章点赞数
    PRIMARY KEY (id)
);

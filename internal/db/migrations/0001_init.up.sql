CREATE TABLE images (
    id           SERIAL PRIMARY KEY,
    img_url      VARCHAR(255) NOT NULL UNIQUE,
    created_at   TIMESTAMPTZ    NOT NULL DEFAULT now(),
    upvotes      INT            NOT NULL DEFAULT 0,
    downvotes    INT            NOT NULL DEFAULT 0,
    total_views  INT            NOT NULL DEFAULT 0,
    caption      VARCHAR(255) ,
);
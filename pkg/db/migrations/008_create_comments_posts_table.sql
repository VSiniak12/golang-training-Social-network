-- +goose Up
create table comments_posts
(
    comment_id serial  not null
        constraint comments_posts_pk
            primary key,
    text       text    not null,
    date_write date    not null,
    user_write integer
        constraint comments_posts_user_id_user_fk
            references users
            on delete set null,
    post_id    integer not null
        constraint comments_posts_posts_id_post_fk
            references posts
            on delete cascade
);

INSERT INTO comments_posts (text, date_write, user_write, post_id) VALUES
('Why??', '2021-04-12', 4, 2),
('is best?', '2021-02-12', 5, 3),
('Hello, how are you?', '2021-04-12', 2, 4),
('Yes, of course', '2021-04-12', 2, 5),
('You are hard-working', '2021-04-10', 3, 5);

-- +goose Down
DROP TABLE comments_posts;
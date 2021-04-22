-- +goose Up
create table posts
(
    id_post   serial  not null
        constraint post_pk
            primary key,
    text      varchar not null,
    date_post date,
    user_id   integer not null
        constraint posts_user_id_user_fk
            references users
            on delete cascade
);

INSERT INTO posts (text, date_post, user_id) VALUES
('My holiday', '2021-04-12', 1),
('My heart is broken', '2021-04-10', 3),
('My mood...', '2021-02-01', 2),
('Hello, friends!', '2021-03-22', 1),
('I will get offer from epam', '2021-04-09', 4);

-- +goose Down
DROP TABLE posts;
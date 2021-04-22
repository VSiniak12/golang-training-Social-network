-- +goose Up
create table messages
(
    id_message   serial  not null
        constraint message_pk
            primary key,
    text         varchar not null,
    from_user_id integer
        constraint messages_user_id_user_fk
            references users
            on delete set null,
    to_user_id   integer
        constraint messages_user_id_user_fk_2
            references users
            on delete set null
);

INSERT INTO messages (text, from_user_id, to_user_id) VALUES
('Hello, how are you?', 1, 2),
('Hello, i am pretty good, and you?', 2, 1),
('Me too', 1, 2),
('Call me, please', 3, 4),
('When do you need an answer?', 4, 1);

-- +goose Down
DROP TABLE messages;
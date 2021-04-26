-- +goose Up
create table friends
(
    id         serial  not null
        constraint friends_pk
            primary key,
    profile_id integer not null
        constraint friends_user_id_user_fk
            references users
            on delete cascade,
    friend_id  integer not null
        constraint friends_user_id_user_fk_2
            references users
            on delete cascade
);

INSERT INTO friends (profile_id, friend_id) VALUES (1,2), (1,4), (3,1), (5,2), (4,3);

-- +goose Down
DROP TABLE friends;
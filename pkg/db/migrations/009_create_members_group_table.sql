-- +goose Up
create table members_group
(
    id       serial  not null
        constraint members_group_pk
            primary key,
    group_id integer not null
        constraint members_group_groups_id_group_fk
            references groups
            on delete cascade,
    user_id  integer not null
        constraint members_group_user_id_user_fk
            references users
            on delete cascade
);

INSERT INTO members_group (group_id, user_id) VALUES (1,2), (1,5), (3,2), (2,1), (5,3);

-- +goose Down
DROP TABLE members_group;
-- +goose Up
create table administrator_group
(
    id               serial  not null
        constraint administrator_group_pk
            primary key,
    group_id         integer not null
        constraint administrator_group_groups_id_group_fk
            references groups
            on delete cascade,
    administrator_id integer not null
        constraint administrator_group_user_id_user_fk
            references users
            on delete cascade
);

INSERT INTO administrator_group (group_id, administrator_id) VALUES (1,2), (2,3), (3,4), (4,1), (5,5);

-- +goose Down
DROP TABLE administrator_group;
-- +goose Up
create table groups
(
    id_group    serial  not null
        constraint groups_pk
            primary key,
    name        varchar not null,
    date_create date    not null
);

INSERT INTO groups (name, date_create) VALUES ('epam', '2021-04-08'),
('Max Korzh', '2020-02-03'), ('Footbal', '2021-01-03'), ('News', '2021-03-21'), ('MDK', '2021-01-30');

-- +goose Down
DROP TABLE groups;
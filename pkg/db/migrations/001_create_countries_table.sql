-- +goose Up
create table countries
(
    country_id serial  not null
        constraint country_pk
            primary key,
    name       varchar not null
);

INSERT INTO countries (name) VALUES ('Belarus'), ('Russia'), ('Ukraina'), ('Czech'), ('Poland');

-- +goose Down
DROP TABLE countries;
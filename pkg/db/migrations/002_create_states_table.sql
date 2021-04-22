-- +goose Up
create table states
(
    id_state   serial  not null
        constraint states_pk
            primary key,
    name       varchar not null,
    country_id integer not null
        constraint states_countries_country_id_fk
            references countries
            on update cascade on delete cascade
);
INSERT INTO states (name, country_id) VALUES
            ('Gomelskaia', 1), ('Minskaaia', 1),('Moskovskaia', 2), ('Kievskaia', 3), ('Xarkovskaia', 3) ;

-- +goose Down
DROP TABLE states;
-- +goose Up
create table cities
(
    id_city  serial  not null
        constraint city_pk
            primary key,
    name     varchar not null,
    state_id integer not null
        constraint "cities _states_state_id_fk"
            references states
            on update cascade on delete cascade
);
INSERT INTO cities (name, state_id) VALUES
('Gomel', 1), ('Mozyr', 1), ('Minsk', 2), ('Moskva', 3), ('Kiev', 4);

-- +goose Down
DROP TABLE cities;
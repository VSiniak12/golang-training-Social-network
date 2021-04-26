-- +goose Up
create table users
(
    id_user    serial  not null
        constraint user_pk
            primary key,
    login      varchar,
    password   varchar,
    gender     boolean not null,
    email      varchar not null,
    last_name  varchar not null,
    first_name varchar not null,
    birthday   date    not null,
    city_id    integer
        constraint user_cities_city_id_fk
            references cities
            on update set null on delete set null
);

INSERT INTO users(login, password, gender, email, last_name, first_name, birthday, city_id) VALUES
('login1', '1234', true, 'e1@mail.ru', 'Ivanov', 'Ivan', '1997-03-01', 1),
('login2', '1234', true, 'e2@mail.ru', 'Sidorov', 'Andrei', '2000-12-03', 1),
('login3', '1234', false, 'e3@mail.ru', 'Sokolova', 'Svetlana', '1990-05-01', 3),
('login4', '1234', false, 'e5@mail.ru', 'Orexova', 'Anna', '2002-01-22', 5),
('login5', '1234', true, 'e5@mail.ru', 'Bulkin', 'Slava', '1980-06-03', 4);

-- +goose Down
DROP TABLE users;

--Insert in tables

--Insert countires
INSERT INTO countries (name) VALUES ('Belarus'), ('Russia'), ('Ukraina'), ('Czech'), ('Poland');

--Insert states
INSERT INTO states (name, country_id) VALUES
('Gomelskaia', 1), ('Minskaaia', 1),('Moskovskaia', 2), ('Kievskaia', 3), ('Xarkovskaia', 3) ;

--Insert cities
INSERT INTO cities (name, state_id) VALUES
('Gomel', 1), ('Mozyr', 1), ('Minsk', 2), ('Moskva', 3), ('Kiev', 4);

--Insert groups
INSERT INTO groups (name, date_create) VALUES ('epam', '2021-04-08'),
('Max Korzh', '2020-02-03'), ('Footbal', '2021-01-03'), ('News', '2021-03-21'), ('MDK', '2021-01-30');

--Insert users
INSERT INTO users(login, password, gender, email, last_name, first_name, birthday, city_id) VALUES
('login1', '1234', true, 'e1@mail.ru', 'Ivanov', 'Ivan', '1997-03-01', 1),
('login2', '1234', true, 'e2@mail.ru', 'Sidorov', 'Andrei', '2000-12-03', 1),
('login3', '1234', false, 'e3@mail.ru', 'Sokolova', 'Svetlana', '1990-05-01', 3),
('login4', '1234', false, 'e5@mail.ru', 'Orexova', 'Anna', '2002-01-22', 5),
('login5', '1234', true, 'e5@mail.ru', 'Bulkin', 'Slava', '1980-06-03', 4);

--Insert posts
INSERT INTO posts (text, date_post, user_id) VALUES
('My holiday', '2021-04-12', 1),
('My heart is broken', '2021-04-10', 3),
('My mood...', '2021-02-01', 2),
('Hello, friends!', '2021-03-22', 1),
('I will get offer from epam', '2021-04-09', 4);

--Insert messages
INSERT INTO messages (text, from_user_id, to_user_id) VALUES
('Hello, how are you?', 1, 2),
('Hello, i am pretty good, and you?', 2, 1),
('Me too', 1, 2),
('Call me, please', 3, 4),
('When do you need an answer?', 4, 1);

--Insert comments_posts
INSERT INTO comments_posts (text, date_write, user_write, post_id) VALUES
('Why??', '2021-04-12', 4, 2),
('is best?', '2021-02-12', 5, 3),
('Hello, how are you?', '2021-04-12', 2, 4),
('Yes, of course', '2021-04-12', 2, 5),
('You are hard-working', '2021-04-10', 3, 5);

--Insert members_group
INSERT INTO members_group (group_id, user_id) VALUES (1,2), (1,5), (3,2), (2,1), (5,3);

--Insert friends
INSERT INTO friends (profile_id, friend_id) VALUES (1,2), (1,4), (3,1), (5,2), (4,3);

--Insert administrator_group
INSERT INTO administrator_group (group_id, administrator_id) VALUES (1,2), (2,3), (3,4), (4,1), (5,5);
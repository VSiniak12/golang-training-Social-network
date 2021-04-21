package data

//users
const readAllUsersQuery = `SELECT * FROM "users"`
const readAllUsersWhereQuery = `SELECT * FROM "users" WHERE (id_user = $1)`
const insertUserQuery = `INSERT INTO "users" ("id_user","login","password","gender","email","last_name","first_name","birthday","city_id") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING "users".*`
const deleteUserQuery = `DELETE FROM "users" WHERE (id_user = $1)`
const readUserQuery = `SELECT * FROM "users" WHERE (id_user = $1) LIMIT 1`

//states
const readAllStatesQuery = `SELECT * FROM "states"`
const readStateQuery = `SELECT * FROM "states" WHERE (id_state = $1) LIMIT 1`
const insertStateQuery = `INSERT INTO "states" ("id_state","name","country_id") VALUES ($1,$2,$3) RETURNING "states".*`
const selectStatesWhereQuery = `SELECT * FROM "states"  WHERE (id_state = $1)`
const deleteStateQuery = `DELETE FROM "states" WHERE (id_state = $1)`

//usersCity
const readAllUserCityQuery = `SELECT users.id_user, users.login, users.password, users.gender,users.email, users.last_name, users.first_name, users.birthday, cities.name FROM "social_network"."users" left join social_network.cities on social_network.users.city_id = social_network.cities.id_city`
const joinNameCity = "left join social_network.cities on social_network.users.city_id = social_network.cities.id_city"

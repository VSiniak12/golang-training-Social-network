## golang-training-Social-network
    
### Description
Social Network Application.  
It contains technologies: golang, gorm, testify/assert, go-sqlmock, gorilla/mux. Database is PostgresSQL

### How to run it
Command: ```go run ./cmd/main.go```  
host: localhost  
port: 8080  

### How to run unit tests
Command: ```go test ./pkg/data/```  

### Table with Paths 

Path | Method | Description | Body  
---| --- | --- | ---
/users/{id:[0-9]+} | GET | get user by id | Response: ```json{"IdUser":12,"Login":"login6","Password":"12345","Gender":false,"Email":"e1@mail.ru","LastName":"Ivanov","FirstName":"Ivan","Birthday":"1997-03-01T00:00:00Z","CityId":3}```
/users/{id:[0-9]+} | DELETE | delete user by id | No body
/users/{id:[0-9]+} | PUT | update user | Request: ```json{"Login":"login6","Password":"12345","Gender":false,"Email":"e1@mail.ru","LastName":"Ivanov","FirstName":"Ivan","Birthday":"1997-03-01T00:00:00Z","CityId":3}```
/users/ | GET | get all users | Response: ```json[{"IdUser":1,"Login":"login1","Password":"1234","Gender":true,"Email":"e1@mail.ru","LastName":"Ivanov","FirstName":"Ivan","Birthday":"1997-03-01T00:00:00Z","CityId":1},{"IdUser":2,"Login":"login2","Password":"1234","Gender":true,"Email":"e2@mail.ru","LastName":"Sidorov","FirstName":"Andrei","Birthday":"2000-12-03T00:00:00Z","CityId":1},{"IdUser":3,"Login":"login3","Password":"1234","Gender":false,"Email":"e3@mail.ru","LastName":"Sokolova","FirstName":"Svetlana","Birthday":"1990-05-01T00:00:00Z","CityId":3},{"IdUser":4,"Login":"login4","Password":"1234","Gender":false,"Email":"e5@mail.ru","LastName":"Orexova","FirstName":"Anna","Birthday":"2002-01-22T00:00:00Z","CityId":5},{"IdUser":5,"Login":"login5","Password":"1234","Gender":true,"Email":"e5@mail.ru","LastName":"Bulkin","FirstName":"Slava","Birthday":"1980-06-03T00:00:00Z","CityId":4}]```
/users/ | POST | create user | Request: ```json{"Login":"login6","Password":"1234","Gender":true,"Email":"e1@mail.ru","LastName":"Ivanov","FirstName":"Ivan","Birthday":"1997-03-01","CityId":2}```
/states/{id:[0-9]+} | GET | get state by id | Response: ```json{"IdState":19,"Name":"Mogilevskaia","CountryId":4}```
/states/{id:[0-9]+} | DELETE | delete state by id | No body
/states/{id:[0-9]+} | PUT | update state | ```json{"name" : "Mogilevskaia", "CountryId" : 4}```
/states/ | GET | get all states | Response: ```json[{"IdState":1,"Name":"Gomelskaia","CountryId":1},{"IdState":2,"Name":"Minskaaia","CountryId":1},{"IdState":3,"Name":"Moskovskaia","CountryId":2},{"IdState":4,"Name":"Kievskaia","CountryId":3},{"IdState":5,"Name":"Xarkovskaia","CountryId":3}]```
/states/ | POST | create state | Request: ```json"name":"Brestskaia", "CountryId" : 1}```
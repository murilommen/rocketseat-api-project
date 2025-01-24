# Rocketseat Golang API

This repository is a case study proposed by Rocketseat that aims to create a simple REST API in Go 
that interacts with a user database in memory. Instead of using a traditional database, it uses a hashmap 
and deals with users using that simpler data structure. 

The challenge is described [here](https://efficient-sloth-d85.notion.site/Desafio-Criando-uma-API-REST-a25c8593dfb54c1790a863a19a355a5e) and instead of going with the standard library, I decided to use `gin` in order to learn about complimentary concepts not treated on their 
Golang course.

## Architecture
I have decided to implement a different approach than what the course provides, trying to follow the ports+adapters architecture (Hexagonal), in order to facilitate switching from an in-memory database to a traditional database Adapter in the future. It is far from perfect, but the base idea is grounded. 

## Endpoints

- POST `/api/users`: Creates a user using the information provided on the request body

- GET `/api/users`: Returns an array of users

- GET `/api/users/:id`: Returns the User object with the specified ID

- DELETE `/api/users/:id`: Removes the user from the database with the specified user id 
	
- PUT `/api/users/:id`: Updates the user with the specified ID and with the request body that complies with the User database. It returns the updated body too.

The request body for the `POST` and `PUT` operations looks like this: 
```json
{
  "FirstName": "the user's first name",
  "LastName": "user's last name",
  "Biography": "user's biography"
}
```
The app will create an unique identifier for the User, which can be used as a path parameter to `GET`, `DELETE` and `PUT` users afterwards.

## Running the project
You will need `go>=1.22.0` to be able to build this project, then you can run `go run cmd/server/main.go` to make the server start listening to requests.

## Next steps
- [ ] Add string validation for the inputs (currently it only checks if id's aren't empty and if the User body is respected)
- [ ] Switch from in-memory to a File-based approach and also to a traditional local database, such as SQLite or Postgres
- [ ] Add a Dockerfile for an easier setup (and deployment)
- [ ] Unit test modules
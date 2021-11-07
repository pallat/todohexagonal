# TodoMVC API

This repository is an API experimental to serve TodoMVC

## Requitements

**GET** /tokenz

to get simple JWT token

**POST** /todos

to create a todo task

```json
{
    "text": "test 1"
}
```

**GET** /todos

to get todo list

**DELETE** /todos

to delete a todo task given an ID

```json
{
    "text": "test 1"
}
```

## Non-Functional Requirements

- use Gin-Gonic framework
- configuration in environment
- gracefully shutting down
- path /x to get git commit
- liveness probe
- readiness probe
- JWT authentications
- rate limited
- Dockerfile
- error logging at crime scene with tracking id

## TodoMVC Frontend of This API

https://github.com/pallat/todowasm

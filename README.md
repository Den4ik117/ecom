# Backend API in Golang

Simple API in Golang

[YouTube tutorial](https://youtu.be/7VLmLOiQ3ck?si=N9xf8H_wSCO3kz2W)

## Technologies

- Golang
- MySQL
- Docker

## Requirements

- You need to install [migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- You need to install command `make`

## Commands

- Use `make up` to start containers
- Use `make down` to down containers
- Use `make r` to restart containers
- Use `make migration <migration name>` to create migration
- Use `make migrate-up` and `make migrate-down` to up and down migrations
- Use `make run` to run application
- Use `make test` to start tests
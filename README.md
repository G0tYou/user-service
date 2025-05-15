# user-service

## Description
written in Go.

## Code Structure
This project has 4 Domain layer :

- Domain Layer, act as model or entity.
- Repository Layer, calls to database, memory, external services/APIs.
- Service Layer, business login happen here.
- Delivery Layer, routers and handlers.

## Owner
Fandi Dachi

## Run
 - Docker
    * change config.json ( database host to host.docker.internal)
    - docker build -t user-service -f .\Dockerfile .
    - docker run -p 8080:8080 user-service
 - Local
    - go run .\cmd\server\main.go
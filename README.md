# Midterm-Exam-EAI
Midterm Exam for EAI Course, Information Systems, Telkom University

# Introduction
This project is a RESTful API for a simple online store. The API is built using go and fiber as the web framework. The database used is PostgreSQL. The API is DDD (Domain Driven Design) based. The API is also containerized using Docker.

# Project Structure
The project is structured as follows:
```
.
├───cmd
└───internal
    ├───app
    ├───config
    ├───controllers
    ├───domain
    ├───dto
    ├───middleware
    ├───repositories
    ├───routes
    └───services
```
- cmd: contains the main.go file which is the entry point of the application
- internal: contains the internal packages of the application
- app: contains the application logic
- config: contains the configuration files
- controllers: contains the controllers
- domain: contains the domain models
- dto: contains the data transfer objects
- middleware: contains the middleware
- repositories: contains the repositories (data access layer)
- routes: contains the routes
- services: contains the services (business logic)

Why i build this project using DDD (Domain Driven Design) Architecture?
- DDD is a software development approach that focuses on the core logic of the application and separates it from the infrastructure and delivery mechanisms.
- Each layer of the application has its own responsibility and is independent of the other layers.
- The application is easier to maintain and extend.
- Easy to test each layer of the application.
- Easy to change frameworks and technologies when needed.

# How to run the project
- Clone the project
- Create a .env file in the root directory of the project and add the following environment variables:
```
DB_HOST=your_host
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=your_database_name
DB_PORT=databse_port
DB_SSLMODE=disable

PORT=your_app_port
```

- Run the following command to run the project:
```bash
$ go run ./cmd/main.go
```
or in windows:
```bash
go run .\cmd\main.go
``` 

# API Documentation
The API documentation can be found in the following link:
https://documenter.getpostman.com/view/16260600/2s93XzyNWT

# Run via Docker
- pull the image from docker hub
```bash
$ docker pull faruqihafiz/uts-eai
```
- run the image
```bash
$ docker run -p 3000:3000 faruqihafiz/uts-eai
```
- open the browser and go to http://localhost:3000
- access the adminer at http://localhost:8080

# Run via Docker Compose
- pull the image from docker hub
```bash
$ docker pull faruqihafiz/uts-eai
```

- run the image
```bash
$ docker-compose up -d
```
- API will be running at http://localhost:3000
- Adminer will be running at http://localhost:8080







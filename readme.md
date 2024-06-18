# TODO List Application

This is a simple todo list application built with Go and Redis. 
It allows users to create, read, update, and delete their tasks.

## Features
- Users can create an account providing their username, age, gender, and email.
- Users can add, view, update, and delete their todos.

## Prerequisites
- Go installed
- Redis installed
- Docker (optional)

## API Endpoints
- `POST /users`: Create a new user
- `GET /users/{username}`: Get user details
- `PUT /users/{username}`: Update user details
- `DELETE /users/{username}`: Delete a user
- `POST /todos/{username}`: Create a new todo
- `GET /todos/{username}`: Get all todos for a user
- `PUT /todos/{username}/{id}`: Update a todo
- `DELETE /todos/{username}/{id}`: Delete a todo

## Docker
You can use Docker to containerize the application. 
1. Build the Docker image: `docker build -t todo-app .`
2. Run the container: `docker run -p 8080:8080 todo-app`

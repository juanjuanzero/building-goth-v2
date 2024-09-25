# Exploring the GOTH Stack:

Building a To Do app

## Stack Details:

- Golang Backend Server
  - use air for speed of development
- Tailwind for UI
- Templ and HTMX for UI interactivity

## Project Structure

- cmd: will contain the application code like main.go
- internal: will contain the internal business logic that will happen in our application, it will have two main subdirectories
  - handlers: handlers that will be bound to the thing
  - services: business logic for services
- components: for the templ files
- assets: for static assets

## Getting Started

Run go mod to create the backend

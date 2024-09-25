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

Run go mod to create the backend.

## Create a simple http server

- create handler for the home route at /
- addRoutes will be used to add routes
- wire it up for the home page

## Create a your first page with templ

## Add tailwind

## Add layout and other templ components

## Add to do service

## Wire up to do service in handlers to handle request to add TODO item

## Expose other update, delete routes for todo, add htmx for interactivity

## iterating on the UI, working with air and tailwind for auto-updates

## persist TODO items, adding a sqlLite instance

## add user login using clerk

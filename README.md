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

- i already have a the templ cli installed on my machine, and have enabled the templ vscode extension
- run `go get github.com/a-h/templ`
- create a layout component that will encapsulate all of the components in and run `templ generate`
- wire up the layout temple to respond to the home route in the handler

## Add tailwind

- I use node 20 to install tailwindcss.
- add .nvmrc for the node 20 version that you have
- `$ npm install -D tailwindcss && npx tailwindcss init` which creates a tailwind.config.js file
- well create a static folder under src to hold all of our static assets like our css files
- create a file called input.css this is where we will write our
- run `npx tailwindcss -i ./src/static/input.css -o ./public/static/output.css --watch` this runs things in watch mode
- in order to serve the css files, you need to create a file server and serve that as well to, create a new handler to handle requests for static files
- a new learning experience, i got the file server to work, we use stripPrefix to remove the static prefix so that the file server , which uses the root will serve that file

## Add layout and other templ components

## Add to do service

## Wire up to do service in handlers to handle request to add TODO item

## Expose other update, delete routes for todo, add htmx for interactivity

## iterating on the UI, working with air and tailwind for auto-updates

## persist TODO items, adding a sqlLite instance

## add user login using clerk

# Bookings and Reservations

This is the repository for my bookings and reservations project.

## Backend Details

- Built in Go version 1.17
- Uses the [chi router](https://github.com/go-chi/chi)
- Uses Alex Edwards [SCS session management](https://github.com/alexedwards/scs)
- Uses [NoSurf](https://github.com/justinas/nosurf)

## Frontend Details

- Uses [Bootstrap 5.1.3](https://getbootstrap.com/docs/5.1/getting-started/introduction/)
- Uses [notie](https://github.com/jaredreich/notie) for clean and simple notification, input, and selection suite
- Uses [sweetalert2](https://sweetalert2.github.io/) for JavaScript popup boxes
- JavaScript

## Local Development

Run the web server with the command:

`go run cmd/web/*.go`

Goto `http://localhost:8080`

### Helpful Commands

`go get <path to package>`: to install new go dependencies

`go mod tidy`: to update the dependency lock file after adding imports in the project

## Project Layout

### cmd

Under `cmd/web` includes the following:

* `main.go` - entry point of the web app
* `middleware.go` - place for code that interacts with a request betweem the client and server
* `routes.go` - place where the router and routes are configured

### pkg

Under `pkg/config` includes the following:

* `config.go` - stores the application-wide configuration

Under `pkg/handlers` includes the following:

* `handlers.go` - defines the repository connection and the logic for each route

Under `pkg/models` includes the following:

* `templatedata.go` - defines the schema for data that gets injected into the HTML templates

Under `pkg/render` includes the following:

* `render.go` - defines how HTML templates are rendered by reading the filesystem or template cache

### templates

* `layout.html` - defines the content shared across the application
* `index.html` - defines the content a user will see on the application's home page
* `about.html` - defines the content a user will see on the application's about page
* `contact.html` - defines the content a user will see on the application's contact page
* `generals.html` - defines the content a user will see on the application's generals page
* `majors.html` - defines the content a user will see on the application's majors page
* `make-reservation.html` - defines the content a user will see on the application's reservation page
* `search-availability.html` - defines the content a user will see on the application's search availability page

### static

Stores all static files such as images, CSS and JavaScript files.

## Endpoints Available

* `/` - home page
* `/about` - about page
* `/contact` - contact page
* `/generals-quarters` - generals page
* `/majors-suite` - majors page
* `/search-availability` - search availability page
* `/make-reservation` - make reservation page

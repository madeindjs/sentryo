# Sentryo

A small REST web API endpoints to test GO lang. The purpose of this repository is to experiment [Go](https://golang.org/) ecosystem. I intentionally avoid complete Frameworks and ORM to build this app. I use only theses libraries:

- [echo](https://echo.labstack.com/) a small HTTP web server (because using Plain Go HTTP library become ugly when application grow)
- [go-sqlite3](https://github.com/mattn/go-sqlite3) to communicate with database
- [glog](https://github.com/golang/glog) to improve logging system

For the moment, I only implement theses REST endpoints:

- vehicles
  - GET `/`  to print this readme file
  - GET `/vehicles/` to list vehicles
  - GET `/vehicles/:id` to show a vehicle
- starships
  - GET `/starships/` to list starships
  - GET `/starships/:id` to show a starship
- peoples
  - GET `/peoples/` to list peoples
  - GET `/peoples/:id` to show a people
  - POST `/peoples/` to create a people
  - PUT `/peoples/:id` to update a people
  - DELETE `/peoples/:id` to destroy a people

I made choice to adapt my code according to the given database schema (which is no good at all  because there are only `TEXT` fields, no primary/foreigns keys, etc..). Also there are a lack consistency about tables name who mix singular and plural names.

## Usage

Theres some examples using cURL:

Create a people:

    $ curl -X POST -H 'Content-Type: application/json' \
           -d '{"id":"42","name":"Toto","gender":"Male"}' \
           http://localhost:1323/peoples/

Find a people:

    $ curl http://localhost:1323/peoples/42

Update a people:

    $ curl -X PUT -H 'Content-Type: application/json' \
           -d '{"name":"tata","gender":"male"}'  \
          http://localhost:1323/peoples/42

Remove a people:

    $ curl -X DELETE http://localhost:1323/peoples/42

## Installation

Simply run `go get` & build project using Go compiler:

    $ go get github.com/madeindjs/sentryo
    $ cd "$GOPATH/github.com/madeindjs/sentryo"
    $ go build main.go
    
> Please note you need to [Install SQlite](https://www.sqlite.org/download.html) on your system

---

### Go further

To go further I recommend theses actions:

- Use [Dep](https://golang.github.io/dep) package manager
- Add some unit test to tests CRUD Models methods
- Correct SQLite database schema to use primary/foreign keys, correct field types, etc..
- Implement [Thread Safe Singleton](http://marcio.io/2015/07/singleton-pattern-in-go/) for database system
- Convert Markdown to HTML on the landing page
- Factorize Starships & Vehicles (who are basically the same thing)
- Respect [JSON:API](https://jsonapi.org/) norm
- Add a caching system

### Links

There some useful resources I saw during development:

- <https://thenewstack.io/building-a-web-server-in-go/>
- <https://tutorialedge.net/golang/creating-restful-api-with-golang/>
- [List of HTTP Codes in Go net/http](https://golang.org/src/net/http/status.go)
- [Singleton in Go](http://marcio.io/2015/07/singleton-pattern-in-go/)
- [Go Database documentation](http://go-database-sql.org/)

# Sentryo

A small REST web API endpoints to test GO lang.

This use [Echo](https://echo.labstack.com/) Go we framework (because using Plain
Go become uggly AF when application grow). The purpose of this repository is to
test Go ecosystem so I intentionally avoid complete ORM librairies.

For the moment, I only implement theses REST endpoints:

There the current list of REST endpoints about **vehicles**, **users**, and
**spaceships**:

- GET </> to print this readme
- GET </vehicles/> to list vehicles in database
- GET </vehicles/:id> to show a vehicle
- GET </starships/> to list starships in database
- GET </starships/:id> to show a starship
- GET </peoples/> to list peoples in database
- GET </peoples/:id> to show a people
- POST </peoples/> to create a peoples

Some concerns about somes choice during development:

- Sentryo use `TEXT` everywhere in database (which is **very** not good). So I
  made choice to adapt my code according to the databse schema.
- `vehicle` contains a spelling mistake which I included in controller / model /
  url names for constencies reasons
- `people` table should be renamed as `users` (at least it should be `peoples`,
  with a plurial mark)

Theres somes examples using cURL:

- Create a people:

~~~bash
$ curl -X POST -H 'Content-Type: application/json' \
       -d '{"id":"42","name":"Toto","gender":"Male"}' \
       http://localhost:1323/peoples/
~~~

- Find a people:

~~~bash
$ curl http://localhost:1323/peoples/42
~~~

- Update a people:

~~~bash
$ curl -X PUT -H 'Content-Type: application/json' \
       -d '{"name":"tata","gender":"male"}'  \
      http://localhost:1323/peoples/42
~~~

## Instalation

Simply install depencies, clone repository and build project using Go compiler:

~~~bash
$ go get github.com/labstack/echo
$ go get github.com/mattn/go-sqlite3
$ git clone https://github.com/madeindjs/sentryo.git
$ cd sentryo
$ go build main.go
$ ./main
~~~

---

### Go further

- Use [Dep](https://golang.github.io/dep) package manager
- Migrate SQLite database to use index, correct field types; foreign keys, etc..
- Implement [Threadsafe Singleton](http://marcio.io/2015/07/singleton-pattern-in-go/)
  for database system
- Create a controller package
- Convert Markdown to HTML
- Factorize Satsrhips & Vehicles (who are basicaly the same thing)

### Links

- <https://thenewstack.io/building-a-web-server-in-go/>
- <https://tutorialedge.net/golang/creating-restful-api-with-golang/>
- [List of HTTP Codes in Go net/http](https://golang.org/src/net/http/status.go)
- [Singleton in Go](http://marcio.io/2015/07/singleton-pattern-in-go/)

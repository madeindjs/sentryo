# Sentryo

A small REST web API endpoints to test GO lang.

This use [Echo](https://echo.labstack.com/) Go we framework (because using Plain Go become uggly AF when application grow).

## Instalation

~~~bash
$ go get github.com/labstack/echo
$ go get github.com/mattn/go-sqlite3
~~~

## Usage

Install Go then:

~~~bash
$ git clone https://github.com/madeindjs/sentryo.git
$ cd sentryo
$ go build main.go
$ ./main
~~~

## Note

- Sentryo use `TEXT`everywhere in database (which is very not good)
- `vehicle` contains a speling mistake which I included in controller / model / url names for constencies reasons

## Go further

- Use [Dep](https://golang.github.io/dep) package manager
- Migrate SQLite database to use index, correct field types; foreign keys, etc..

## Links

- <https://thenewstack.io/building-a-web-server-in-go/>
- <https://tutorialedge.net/golang/creating-restful-api-with-golang/>
- [List of HTTP Codes in Go net/http](https://golang.org/src/net/http/status.go)
- [Singleton in Go](http://marcio.io/2015/07/singleton-pattern-in-go/)

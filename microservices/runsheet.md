# Microservices

----

Fetch the Repo (locally)

```
$ git clone https://github.com/ContainerSolutions/yow-workshop.git
```

----

## Pre-defined API

Swagger specification:

`./microservices/resources/deals-swagger.json`

----

## Building a new Microservice (v1)

Example Go implementation

`open ./microservices/v1/main.go`

----

## Build and run service

```bash
$ go build .
$ go run main.go &
```

----

## Call service

```bash
$ curl localhost:8080/deals?id=1
```

or

`http://[IP_of_VM]/deals?id=1` from browser

----

## Stop service

```bash
$ kill %1
```

----

- Improve GoLang version
- Use Swagger Code Gen e.g.:
```bash
java -jar /home/swagger-codegen-cli.jar generate \
  -i ./resources/deals-swagger.json \
  -l nodejs-server \
  -o nodejs
```
- Write your own

----


## Review

* Questions?
* On to containers...

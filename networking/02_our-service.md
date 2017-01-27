## Exercise

- Add database (mongo) to service
- Run mongo container
- Run our service image and connect to mongo container

----

## Add a Database (v2)

`./microserices/v2/main.go`

----

Our service expects a `deals-db` to connect to.

We want to:

* Create a network for our two containers
* Run a `mongo` container (hint: container name == hostname)

----

## Possible solution

```bash
cd v2/
docker build -t adobe-workshop/deals:v2 .
docker network create my_network
docker run --name deals-db -d --network my_network mongo
docker run -d -p 8080:8080 --network my_network adobe-workshop/deals:v2
```

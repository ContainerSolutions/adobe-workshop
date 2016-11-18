# Microservices Tooling

---

## Install Weave Scope

```bash
sudo wget -O /usr/local/bin/scope https://git.io/scope
sudo chmod a+x /usr/local/bin/scope
sudo scope launch
```

---

## Launch container with Weave plugin

```bash
eval $(weave env)
docker run -d --publish 8080:8080 --name deals-dev microservices-demo/deals
```

---

## Open in browser

`localhost:4040` 
or
`{docker-machine ip}:4040`

---

## Review

* Questions?
* Thanks for coming!
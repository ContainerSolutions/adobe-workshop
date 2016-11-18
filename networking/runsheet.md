# Container Networking

---

## Install Weave Net

```bash
sudo curl -L git.io/weave -o /usr/local/bin/weave
sudo chmod +x /usr/local/bin/weave
weave launch
```

---

## Launch container with Weave plugin

```bash
docker run -dit $(weave dns-args) --publish 8080:8080 --net=weave -h deals-dev.weave.local microservices-demo/deals
```

---

## Ping container via weave

```bash
docker run -it $(weave dns-args) -h pinger.weave.local --net=weave amouat/network-utils ping -c 1 deals-dev
```

---

## Review

* Questions?
* [On to visualisation...](../visualisation/runsheet.md)

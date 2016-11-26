# Container Networking

---
## Install Docker Machine (& VirtualBox)
```bash
sudo su -
curl -L https://github.com/docker/machine/releases/download/v0.8.2/docker-machine-`uname -s`-`uname -m` >/usr/local/bin/docker-machine
chmod +x /usr/local/bin/docker-machine
sudo apt-get install virtualbox 
```

---

## Install Weave Net
Let's start by fetching the weavenet binary and adjusting the permissions
```bash
sudo curl -L git.io/weave -o /usr/local/bin/weave
sudo chmod +x /usr/local/bin/weave
```

---

## Launch Weave Net
For this scenario we will use Docker Machine as this allows to create multiple hosts.
Let's create our first host (VM) and launch weave net on it

```bash
docker-machine create -d virtualbox host1
eval $(docker-machine env host1)
weave launch
```

This has launched the weave router and the plugin.

---

## Ping my container
As a basic first setup, we'll launch a small (alpine) container and test that we can connect to it from a second container.
Using the --net option we specify that we want to use the weave docker plugin

```bash
docker run -dit $(weave dns-args) --net=weave -h pingme.weave.local alpine sh
docker run -it $(weave dns-args) --net=weave -h pinger.weave.local amouat/network-utils ping -c 1 pingme
```

Notice that we are able to find the first container simply using the hostname.

---

## Multihost setup
Now for a less trivial example, let's setup a second host (VM) again using Docker Machine.
Notice that when we launch weave on the second host we are pointing it to the first host's IP. If we continue to add more hosts, we only ever need to point weave to one other host. The hosts will discover each other via the Gossip protocol.

```bash
docker-machine create -d virtualbox host2
eval $(docker-machine env host2)
weave launch $(docker-machine ip host1)
docker run -it $(weave dns-args) --net=weave -h pinger.weave.local amouat/network-utils ping -c 1 pingme
```

Voila! Multi host networking! These two VMs on our local laptop could just as easily be physical servers on opposite sides of the world.

---

## One step further
Now let's add another instance of our 'pingme' service on host1. We then return to host2 and once more try to ping our service.

```bash
eval $(docker-machine env host1)
docker run -dit $(weave dns-args) --net=weave -h pingme.weave.local alpine sh
eval $(docker-machine env host2)        
docker run -it $(weave dns-args) --net=weave -h pinger.weave.local amouat/network-utils ping -c 1 pingme
docker run -it $(weave dns-args) --net=weave -h pinger.weave.local amouat/network-utils ping -c 1 pingme
...
```
As you run the ping service a few times, we can see that we get responses from both instances. Built-in load balancing!

---

---

## Example using Docker Compose
Sock Shop - Docker single
```bash
git clone https://github.com/microservices-demo/microservices-demo.git
cd microservices-demo/deploy/docker-single
weave launch
docker-compose up -d
```
This will launch the full Sock Shop demo locally. Container communicatin provided by Weave Net plugin.

---

## Launch our own microservice container with Weave plugin

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

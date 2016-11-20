## What is ingress?

Typically, services and pods have IPs only routable by the cluster network. All traffic that ends up at an edge router is either dropped or forwarded elsewhere. Conceptually, this might look like:
```
    internet
        |
  ------------
  [ Services ]
```
An Ingress is a collection of rules that allow inbound connections to reach the cluster services.
```
    internet
        |
   [ Ingress ]
   --|-----|--
   [ Services ]
```

----

It can be configured:
* to give services externally-reachable urls
* load balance traffic
* terminate SSL
* offer name based virtual hosting 

An Ingress controller is responsible for fulfilling the Ingress, usually with a loadbalancer, though it may also configure your edge router or additional frontends to help handle the traffic in an HA manner.

----

### Ingress controller

In order for the Ingress resource to work, the cluster must have an Ingress controller running

An Ingress Controller is a daemon, deployed as a Kubernetes Pod, that watches the ApiServer's /ingresses endpoint for updates to the Ingress resource. Its job is to satisfy requests for ingress.

Workflow:
* Poll until apiserver reports a new Ingress
* Write the LB config file based on a go text/template
* Reload LB config

----

### Example
Ingress resource
```
apiVersion: extensions/v1beta1
kind: Ingress
metadata: 
  name: frontend-ingress
spec: 
  rules: 
    - 
      host: socks.example.com
      http: 
        paths: 
          - 
            backend: 
              serviceName: front-end
              servicePort: 80
            path: /
```
*POSTing this to the API server will have no effect if you have not configured an Ingress controller.*

----

```
apiVersion: v1
kind: ReplicationController
metadata:
  name: traefik-ingress-controller
  labels:
    k8s-app: traefik-ingress-lb
spec:
  replicas: 1
  selector:
    k8s-app: traefik-ingress-lb
  template:
    metadata:
      labels:
        k8s-app: traefik-ingress-lb
        name: traefik-ingress-lb
```

----

```
    spec:
      terminationGracePeriodSeconds: 60
      containers:
      - image: traefik
        name: traefik-ingress-lb
        imagePullPolicy: Always
        ports:
        - containerPort: 80
          hostPort: 80
        - containerPort: 443
          hostPort: 443
        - containerPort: 8080
        args:
        - --web
        - --kubernetes
```

----

### Try it!

* Deploy the sock-shop 'wholeWeaveDemo.yaml'
* Deploy the 'ingress.yaml'
* Deploy the 'traefik-rc.yaml'

Test it
```bash
curl -I -H 'Host: socks.example.com' $(minikube ip)
```

----
## Creating and Managing Deployments

Deployments sit on top of ReplicaSets and add the ability to define how updates to Pods should be rolled out.

In this lab we will combine everything we learned about Pods and Services

----

### Creating Deployments

```
kubectl create -f mini-demo.yaml --validate=false
```

This file consists of multiple deployments, it ensures just that we don't forget something to launch.

----

### Scaling Deployments

Behind the scenes Deployments manage ReplicaSets. Each deployment is mapped to one active ReplicaSet. Use the `kubectl get replicasets` command to view the current set of replicas.

```
kubectl get replicasets
accounts-187022518        1         1         1m
accounts-db-2169092480    1         1         1m
cart-545437588            1         1         1m
cart-db-3629037076        1         1         1m
catalogue-3448421813      1         1         1m
front-end-1198077563      1         1         1m
login-1839577313          1         1         1m
orders-743750211          1         1         1m
orders-db-2527573982      1         1         1m
payment-2162342800        1         1         1m
queue-master-3195257136   1         1         1m
rabbitmq-2202712352       1         1         1m
shipping-746110140        1         1         1m

```

----

### Scaling Deployments

ReplicaSets are scaled through the Deployment or independently. Use the `kubectl scale` command to scale:

```
kubectl scale --replicas=3 rs/front-end-1198077563
replicaset "front-end-1198077563" scaled
```

```
kubectl describe rs front-end-1198077563
```
```
kubectl scale deployments front-end --replicas=2
deployment "front-end" scaled
```
```
kubectl describe deployment fron-end
```
```
kubectl get pods
```
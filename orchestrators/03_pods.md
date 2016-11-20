## Creating and managing pods

At the core of Kubernetes is the Pod. Pods represent a logical application and hold a collection of one or more containers within the same namespace. In this lab you will learn how to:

* Write a Pod configuration file
* Create and inspect Pods
* Interact with Pods remotely using kubectl

We'll create a Pod named `k8s-hello-world` and interact with it using the kubectl.

----

### Creating Pods

Explore the `k8s-hello-world` pod configuration file:

```
cat pod.yaml
apiVersion: v1
kind: Pod
metadata:
  name: k8s-hello-world
  labels:
    app: k8s-hello-world
spec:
  containers:
    - name: k8s-hello-world
      image: muellermich/nodejs-hello
      ports:
        - containerPort: 8080
```
Create the pod using kubectl:

```
kubectl create -f pod.yaml
```

----

### View Pod details

Use the `kubectl get` and `kubect describe` commands to view details for the `k8s-hello-world` Pod:

```
kubectl get pods
```

```
kubectl describe pods <pod-name>
```

----

### Quiz

* What is the IP address of the `k8s-hello-world` Pod?
* What node is the `k8s-hello-world` Pod running on?
* What containers are running in the `k8s-hello-world` Pod?
* What are the labels attached to the `k8s-hello-world` Pod?
* What arguments are set on the `k8s-hello-world` container?

----

### Interact with a Pod remotely

Pods are allocated a private IP address by default and cannot be reached outside of the cluster. Use the `kubectl port-forward` command to map a local port to a port inside the `k8s-hello-world` pod.

Use two terminals. One to run the `kubectl port-forward` command, and the other to issue `curl` commands.

```
kubectl port-forward k8s-hello-world 10080:8080
```

```
curl http://127.0.0.1:10080
```

----

### View the logs of a Pod

Use the `kubectl logs` command to view the logs for the `k8s-hello-world` Pod:

```
kubectl logs k8s-hello-world
```

> Use the -f flag and observe what happens.

----

### Run an interactive shell inside a Pod

Use the `kubectl exec` command to run an interactive shell inside the `k8s-hello-world` Pod:

```
kubectl exec monolith --stdin --tty -c k8s-hello-world /bin/sh
```
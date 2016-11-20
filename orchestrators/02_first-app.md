### Step 1 kubectl basics

The format of a kubectl command is: kubectl [action] [resource]  
This performs the specified action  (like create, describe) on the specified resource (like node, container). You can use --help after the command to get additional info about possible parameters (kubectl get nodes --help).


Check that kubectl is configured to talk to your cluster, by running the kubectl version command:
```bash
kubectl version
```

You can see both the client and the server versions: 1.4.

----

To view the nodes in the cluster, run the kubectl get nodes command:
```bash	
kubectl get nodes
NAME        STATUS    AGE
127.0.0.1   Ready     48m
```

Here we see the available nodes, just one in our case. Kubernetes will choose where to deploy our application based on the available Node resources.

----

### Step 2 deploy our app 

Let’s run our first app on Kubernetes with the kubectl run command. The run command creates a new deployment.

```bash
kubectl run hellonode --image=docker.io/jocatalin/hellonode:v1 --port=8080 

deployment "hellonode" created
```

This performed a few things for you:
* searched for a suitable node
* scheduled the application to run on that Node
* configured the cluster to reschedule the instance on a new Node when needed 

----

### list your deployments

```bash
kubectl get deployments
NAME        DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
hellonode   1         1         1            1           31s
````

We see that there is 1 deployment running a single instance of your app. 

----

### Step 3 View our app

By default deployed applications are visible only inside the Kubernetes cluster. To view that the application output without exposing it externally, we’ll create a route between our terminal and the Kubernetes cluster using a proxy:
```bash
kubectl proxy &
```
We now have a connection between our host and the Kubernetes cluster.

----

### Accessing the application

To see the output of our application, run a curl request in a new terminal window:
```bash
export POD_NAME=$(kubectl get pods -o go-template --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')
echo $POD_NAME
curl http://localhost:8001/api/v1/proxy/namespaces/default/pods/$POD_NAME/
```

The proxy we started enables direct access to the API. The url is the route to the API of the Pod.
In this tutorial we deployed our first Kubernetes application, and then observed that it was running.

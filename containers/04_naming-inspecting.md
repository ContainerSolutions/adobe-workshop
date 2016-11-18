## Naming and inspecting containers
In this lesson, we will learn about an important Docker concept: container naming.
Naming allows us to:
* Reference easily a container.
* Ensure unicity of a specific container.

We will also see the inspect command, which gives a lot of details about a container.

----

### Naming our containers
So far, we have referenced containers with their ID.
We have copy-pasted the ID, or used a shortened prefix.
But each container can also be referenced by its name.
If a container is named prod-db, I can do:
```bash
docker logs prod-db
docker stop prod-db
etc.
```

----

### Default names

When we create a container, if we don't give a specific name, Docker will pick one for us.
It will be the concatenation of:
* A mood (furious, goofy, suspicious, boring...)
* The name of a famous inventor (tesla, darwin, wozniak...)

Examples: happy_curie, clever_hopper, jovial_lovelace ...

----

### Specifying a name
Specifying a name
You can set the name of the container when you create it.
```bash
docker run --name ticktock jpetazzo/clock
```

If you specify a name that already exists, Docker will refuse to create the container.

This lets us enforce unicity of a given resource.

----

### Renaming containers
Since Docker 1.5 (released February 2015), you can rename containers with docker rename.   
This allows you to "free up" a name without destroying the associated container, for instance.

----

### Inspecting a container

The `docker inspect` command will output a very detailed JSON map.

```JSON
[
    {
        "Id": "4d7820d373a1b408bc1defb9f8ec5e7e01b898b29ba2ee2a4d2b6f2107119b0c",
        "Created": "2016-08-16T13:05:38.672627589Z",
        "Path": "/bin/sh",
        "Args": [
            "-c",
            "/hello"
        ],
        "State": {
            "Status": "exited",
            "Running": false,
            "Paused": false,
            "Restarting": false,
            "OOMKilled": false,
            "Dead": false,
            "Pid": 0,
            "ExitCode": 0,
            "Error": "",
            "StartedAt": "2016-08-16T13:05:39.082040864Z",
            "FinishedAt": "2016-08-16T13:05:39.214338017Z"
```

----

### Parsing JSON with the shell

You could grep and cut or awk the output of docker inspect.
* But it`s a PITA
* If you really must parse JSON from the Shell, use JQ!
```bash
docker inspect <containerID> | jq .
```

We will see a better solution which doesn't require extra tools.

----

### Using --format
You can specify a format string, which will be parsed by Go's text/template package.
```bash
docker inspect --format '{{ json .Created }}' ubuntu_c
"2016-07-25T13:16:10.60041834Z"
```

* The generic syntax is to wrap the expression with double curly braces.
* The expression starts with a dot representing the JSON object.
* Then each field or member can be accessed in dotted notation syntax.
* The optional json keyword asks for valid JSON output.  
(e.g. here it adds the surrounding double-quotes.)
## Advanced Dockering

For elite dockerizers only

----

```bash
$ docker run -it --rm jess/hollywood
```

credit @jessfrazz

----

```bash
docker run -it supertest2014/nyan
```

----

```bash
docker run -t -i -d -p 25565:25565 \
-v /var/run/docker.sock:/var/run/docker.sock \
--name dockercraft \
gaetan/dockercraft
```

----

```bash
docker run -it \
    -v /tmp/.X11-unix:/tmp/.X11-unix \ # mount the X11 socket
    -e DISPLAY=unix$DISPLAY \ # pass the display
    --name cathode \
    jess/1995

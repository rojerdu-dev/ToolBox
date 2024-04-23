Dockerfiles

Create Dockerfiles

To build Docker images, you first need a Dockerfile. A Dockerfile is a text file named Dockerfile and has no file extension.

```
# syntax=docker/dockerfile:1
FROM <base image>  

WORKDIR <path>

COPY <host-path> <image-path>

RUN <command>

ENV <name> <value>

EXPOSE <port-number>

USER <user-or-uid>

CMD ["<command>", "<arg1>"]
 

```

Example Dockerfile for Go
```
# syntax=docker/dockerfile:1
FROM golang:1.22.2-bullseye

# Set destination for COPY 
WORKDIR /app

# Download Go Modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o ./api

# Optional - document in Dockerfile what ports app will listen on by default
EXPOSE 8080

# Run
CMD ["./api"]

```

Create Multi-stage Dockerfiles
A multi-stage build build can carry over artifacts from one build stage to another. Every build stage can be instantiated from a different base image. The resulting image is much lighter as a result.

```
# syntax=docker/dockerfile:1
FROM golang:1.22.2-bullseye AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ./api

# Run tests in container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage ./api ./api

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["./api"]


```

Create Docker Images

Use ```docker build <path>``` to create a Docker image. ```<path>``` sets the context for the docker build command.


1. Use the CWD of the host machine to find the Dockerfile

```$ docker build .                                 ```

2. Specify Dockerfile (-f, --file)

```$ docker build --file ./path/to/Dockerfile       ```


3. Tag an image (-t, --tag)

```$ docker build -t . <user>/<repo>:<version>      ```

```$ docker build -t . rhaegarrr/test:0.1.0         ```

4. List Images
```$ docker image ls                                ```


Create Docker Containers
Docker containers are just running images

Run Docker container from image
```$ docker run <image-name>                                ```

Run public images from Docker Hub or private repo
```$ docker run https://privateregistroy.com/<image-name>   ```

Run in detached mode
```$ docker run --detached <image-name>                     ```

List all containers
```$ docker container ls                                    ```
```$ docker ps                                              ```

List all containers, including stopped ones
```$ docker container ls --all                              ```
```$ docker ps -a                                           ```

Rename a container
```$ docker rename curr_container_name new_container_name   ```

Stop a container
```$ docker stop <container-id>                             ```

Remove a container, but this only works on stopped containers
```$ docker rm <container-id>                               ```

Start a stopped container
```$ docker start <container-id>                            ```

Restart a running container
```$ docker restart <container-id>                          ```

Automatically remove a container when it's stopped
```$ docker run -rm <image-name>                            ```


Access Docker Containers
In Docker, publishing a container's ports means it becomes availabe to not only the Docker host but also the outside world as well.
```$ docker run --publish <host-port>:<container-port> <image-name> ```

Execute shell command in a container
```$ docker exec --interactive --tty <container-id> sh```
```$ docker exec -it --tty <container-id> /bin/bash```


Create Docker Volumes
Volumes are needed to persist data from Docker containers between starts. When a container is removed, all data from the container will be lost if you do not use volumes.

Use a named volume (Docker handles location on host)
```$ docker run --volume <volume-name>:/path/in/container <image-name>   ```

Use a mounted volume (You handle location on host)
```$ docker run --volume /path/on/host:/path/in/container <image-name>   ```

List all volumes, including metadata
```$ docker volume ls                                                    ```
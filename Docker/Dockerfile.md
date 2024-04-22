TODO

To build Docker images, you must first start with a Dockerfile.

A Dockerfile is a text file named Dockerfile and has no file extension

Create Dockerfiles

```
FROM <base image> # get from https://hub.docker.com

# set working directory inside image
WORKDIR /app

# copy source code from host machine to image
# (host: your development machine or server)
# the first dot (.) is the current working directory (CWD)
# of the host. 
# the second dot (.) is the CWD of the image
COPY . .

# run any shell command inside the image
RUN <command> # e.g. [npm run build]

# does not really expose the port 80 
# serves only as documentation so we know which port to expose
EXPOSE 80   # expose ports with the --publish flag

# command gets executed when container starts
CMD <command>  # e.g. ["node", "server.js"]

```

Create Multi-Stage Dockerfiles

Multistage Dockerfiles ar eused to optimize Dockerfiles. 
One use case is to create a builder and a serve stage with separate base images. 
This strategy can be used to make the final image smaller and have a lower attack because it has fewer system libraries.
Each stage starts with FROM.

```
# As lets you alias the current stage with a variable name
FROM <base image> as builder

# now do something, install dependencies, or build your code
RUN <command>  # e.g., g++ main.cpp

# second stage can use a smaller image
# small images are based on alpine 
# or you can build FROM scratch (this is a Docker image)
# if you do not need any system libraries
FROM scratch as serve

# now copy files from builder stage
# e.g. copy binary file that you build in that stage
COPY --from=builder ./api ./api

# command gets executed when container starts
CMD ["./api"]

```

Create Docker Images
```
# <path> sets the context for the docker build command
# use dot (.) to use the CWD of the host machine to find the Dockerfile
$ docker build <path>
$ docker build .

 $ docker build --file ./path/to/Dockerfile
```
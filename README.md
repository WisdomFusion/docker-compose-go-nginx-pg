# docker-compose-go-nginx-pg

Docker Compose for Golang Server with nginx Proxy and PostgreSQL(pg) Database

## Before Usage

With this suite of docker orchestration files, I put a prefix before containers' name, `myapp` be default. In this way, all of the containers built with the compose file would come up this a `myapp-` in its title, for example, `myapp-api` and `myapp-db`, etc.

Copy `.env.example` to `.env`, and change environment keys in `.env` file:

- APP_NAME

## Start Up

With all of your services defined in your `docker-compose` file, you just need to issue a single command to start all of the containers, create the volumes, and set up and connect the networks:

```sh
docker-compose up -d
```

When you run docker-compose up for the first time, it will download all of the necessary Docker images, which might take a while. Once the images are downloaded and stored in your local machine, Compose will create your containers. The `-d` flag daemonizes the process, running your containers in the background.

Once the process is complete, use the following command to list all of the running containers:

```sh
docker ps
```

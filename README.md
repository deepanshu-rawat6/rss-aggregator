# RSS-Aggregator

## Postgres Setup

### Creating volume

For persistent data, to be used in our DB, we can use this command:

```docker
docker volume create postgresqldata
```

### Installing postgres using Docker

```docker
docker run -d -v postgresqldata:/data/db -e POSTGRES_PASSWORD=postgres --name postgres -p 5432:5432 postgres
```

### Installing pdAdmin using Docker

```docker
docker run --name pgadmin -e "PGADMIN_DEFAULT_EMAIL=name@example.com" -e "PGADMIN_DEFAULT_PASSWORD=admin" -p 5050:80 -d dpage/pgadmin4 
```

### Creating a seperate network 

```docker
docker network create --driver bridge pgnetwork
```

Connecting `postgres` and `pgadmin` with the created network

```docker
docker network connect pgnetwork pgadmin
docker network connect pgnetwork postgres
```


## SQLC Setup

Command to use SQLC (docker image)

```docker
docker run --rm -v "$(pwd):/src" -w /src kjconroy/sqlc generate
```
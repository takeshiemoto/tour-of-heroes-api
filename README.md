# tour-of-heroes-api

## Database

```
docker-compose up
docker exec -i -t tour-of-heroes-api_postgres_1 bash
```

```
su - postgres
psql
```

```
CREATE ROLE toh LOGIN CREATEDB PASSWORD 'toh';
CREATE DATABASE toh OWNER toh;
```

```
psql --username=toh --password --dbname=toh
```

```postgresql
create table heroes (
    id serial primary key ,
    name varchar(255),
    createdAt timestamp not null default current_timestamp,
    updateAt timestamp not null default current_timestamp
)
```


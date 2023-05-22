# echo-fw
Build REST API in GO using Echo and Clean Architecture


This is an example how i use Echo to build a REST API

Stack that i use:
- Echo
- PostgreSQL
- Redis

## Create User Table
```
CREATE TABLE IF NOT EXISTS users (
            id serial primary key,
            username varchar(100) NOT NULL,
            first_name varchar(100),
            last_name varchar(100),
            email varchar(100) NOT NULL,
            phone_number varchar(100),
            location varchar(100),
            password varchar(100) NOT NULL,
            salt varchar(100) NOT NULL,
            created_at timestamp,
            updated_at timestamp
        )
```
## Create Posts Table
```
 CREATE TABLE IF NOT EXISTS posts (
            id serial primary key,
            user_id int NOT NULL,
            body text NOT NULL,
            created_at timestamp,
            deleted_at timestamp,
            CONSTRAINT fk_posts
            FOREIGN KEY(user_id)
            REFERENCES users(id)
        )
```
Or you can run this script to create both tables but you need to edit the env files.
```
go run cmd/api/main.go
```

## Postman Documentation
[![Postman](https://cdn.iconscout.com/icon/free/png-512/free-postman-3521648-2945092.png?f=avif&w=32)](https://documenter.getpostman.com/view/13820554/2s93eZzBrj)
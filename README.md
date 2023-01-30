# go-cqrs-project

## Description

The whole structure is based on clean architecture. packaging is by tech. Although you preferable one is packaging by
feature.
the reason for choosing the packaging is its is more understandable for none-golang community which I guess you prefer.
I do not use any web framework but rather used some libraries to be more comfortable. I implement event-bus and used it
with
combination of queue which is implemented by redis. However, to be more convenient I just chose redis as a key value
database.
But maybe it is not perfect choice in this context. As a matter of Clean Code & Architecture, OOP, SOLID principles
there are many situations in code that can be revised but for the sake of time limitation deferred to the future.


## Usage
### Docker compose
to run project by docker please go `docker-compatible-version` branch. from [here](https://github.com/mjedari/go-cqrs-template/tree/docker-compatible-version)

### Make
You can run the application by make file in `src` directory but before make sure that your `redis-server` is running
and you have installed `go` on your machine:
```
make start
```

## Api Guide

### create user (post method)

```
http://localhost:8080/user/create
```

```json
{
  "name": "Hadi",
  "balance": 1000
}
```

### get user(s) (get method)

```
http://localhost:8080/user/{user_id}
http://localhost:8080/user/all
```

### create coin (post method)

```
http://localhost:8080/coin/create
```

```json
{
  "name": "ABAN",
  "price": 3,
  "min": 10
}
```

### get coin(s) (get method)

```
http://localhost:8080/coin/{coin_id}
http://localhost:8080/coin/all
```

### create user (post method)

```
http://localhost:8080/order/create
```

```json
{
  "user_id": 1,
  "coin_id": 1,
  "quantity": 2
}
```

### get order(s) (get method)

```
http://localhost:8080/order/{order_id}
http://localhost:8080/order/all
```



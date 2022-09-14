# Go training

1. Create database
```sh
docker run --name mysql -d -p 3306:3306 -e ALLOW_EMPTY_PASSWORD=yes -e MYSQL_USER=food_delivery -e MYSQL_PASSWORD=anhyeuem -e MYSQL_DATABASE=food_delivery bitnami/mysql:5.7
```

2. Env

```sh
GIN_MODE=debug
MYSQL_CONN_STRING=food_delivery:anhyeuem@tcp(127.0.0.1:3306)/food_delivery?charset\=utf8mb4&parseTime\=True&loc\=Local
```


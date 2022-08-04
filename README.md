# ddd-go
## This project use the recommendation from [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

#### This is an API to study and apply knowledge about DDD

### Requirements
* [Go](https://golang.org/doc/install) >= 1.16
* [Docker](https://docs.docker.com/get-docker/)
* [Docker-compose](https://docs.docker.com/compose/)
* [Postman](https://www.postman.com/downloads/) <b><-Import postman collection from /scripts/go-ddd.postman_collection.json</b>

### See all help commands
```
make help
```

### Run api
```
make copy-env
make api-up
```

### Project tree
````
├── Makefile
├── README.md
├── configs
│   └── container.go
├── coverage.out
├── deployments
│   └── docker-compose.yml
├── go.mod
├── go.sum
├── internal
│   ├── api
│   │   ├── api.go
│   │   ├── handlers
│   │   │   ├── customers.go
│   │   │   └── handlers.go
│   │   └── routes.go
│   ├── domain
│   │   └── customers
│   │       ├── customers.go
│   │       └── interface.go
│   ├── entity
│   │   └── entity.go
│   └── repository
│       ├── customers.go
│       ├── customers_test.go
│       ├── interface.go
│       ├── main_test.go
│       ├── repository.go
│       └── sql.go
├── main.go
└── scripts
    ├── go-ddd.postman_collection.json
    └── migrations
        ├── 0001.down.sql
        └── 0001.up.sql
````

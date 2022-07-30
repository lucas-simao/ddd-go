# ddd-go

#### This is an API to study and apply knowledge about ddd

### Requirements
* [Go](https://golang.org/doc/install) >= 1.16
* [Docker](https://docs.docker.com/get-docker/)
* [Docker-compose](https://docs.docker.com/compose/)

### To run
1 - make copy-env
2 - make api-up

### Project tree
````
├── Makefile
├── README.md
├── configs
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
│       ├── interface.go
│       ├── repository.go
│       └── sql.go
├── main.go
└── scripts
    ├── migrate.sh
    └── migrations
        ├── 0001.down.sql
        └── 0001.up.sql

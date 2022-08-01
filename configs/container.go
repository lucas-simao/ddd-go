package configs

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
)

var (
	user       = "postgres"
	password   = "secret"
	database   = "postgres"
	port       = "5532"
	dataSource = "postgres://%s:%s@localhost:%s/%s?sslmode=disable"
)

type Container struct {
	DB       *sqlx.DB
	pool     *dockertest.Pool
	resource *dockertest.Resource
}

func ContainerRun() *Container {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "13",
		Env: []string{
			"POSTGRES_USER=" + user,
			"POSTGRES_PASSWORD=" + password,
			"POSTGRES_DB=" + database,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: port},
			},
		},
	}

	resource, err := pool.RunWithOptions(&opts, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err.Error())
	}

	time.Sleep(10 * time.Second)

	var postgresUrl = fmt.Sprintf(dataSource, user, password, port, database)

	db, err := sqlx.Connect("postgres", postgresUrl)
	if err != nil {
		log.Panic(err)
		return &Container{}
	}

	err = resource.Expire(2 * 60)
	if err != nil {
		log.Print(err)
	}

	os.Setenv("DATABASE_URL", postgresUrl)

	return &Container{
		db, pool, resource,
	}
}

func (c *Container) ContainerDown() {
	if err := c.pool.Purge(c.resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func (c *Container) RunMigrations(migrationsDir string) {
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		log.Fatalf("Error read migrations dir: %v", err)
	}

	var extensionNameUpMigrations = ".up.sql"

	for i := range files {
		if strings.Contains(files[i].Name(), extensionNameUpMigrations) {
			fileData, err := ioutil.ReadFile(fmt.Sprintf("%v/%v", migrationsDir, files[i].Name()))
			if err != nil {
				log.Fatalf("Error read file: %s, %v", files[i].Name(), err)
			}
			_, err = c.DB.Exec(string(fileData))
			if err != nil {
				log.Fatalf("Error run migrations: %v", err)
			}
		}
	}
}

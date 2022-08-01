package repository

import (
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/lucas-simao/ddd-go/configs"
)

var (
	repo Repository
	DB   *sqlx.DB
)

func TestMain(m *testing.M) {
	newContainer := configs.ContainerRun()
	newContainer.RunMigrations("../../scripts/migrations")
	repo = New()
	DB = newContainer.DB

	code := m.Run()
	newContainer.ContainerDown()
	os.Exit(code)
}

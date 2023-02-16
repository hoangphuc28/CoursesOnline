package utils

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/hoangphuc28/CoursesOnline/API-Gateway/config"
)

func RunDBMigration(c *config.Config) {
	dbURL := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s",
		c.Mysql.User,
		c.Mysql.Password,
		c.Mysql.Host,
		c.Mysql.Port,
		c.Mysql.DBName,
	)

	migration, err := migrate.New(c.Services.MigrationURL, dbURL)
	if err != nil {
		fmt.Println("Cannot create migration instance", err)
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println("Cannot run migrate up", err)
	}
}

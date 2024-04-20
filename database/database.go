package database

import (
	"database/sql"
	"fmt"

	"embed"

	migrate "github.com/rubenv/sql-migrate"
)

var (
	DbConnection *sql.DB
)

var dbMigrations embed.FS

func DbMigrate(dbParam *sql.DB) {
	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "sql_migrations",
	}
	n, errs := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	if errs != nil {
		panic(errs)
	}

	DbConnection = dbParam

	fmt.Println("Apllied", n, "migrations!")
}

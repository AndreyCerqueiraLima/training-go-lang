package migrations

// import (
// 	"context"
// 	"os"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func Run(ctx context.Context) {
// 	db := env.FromContext(ctx).DB
// 	migrationDriver, err := migrateSql.WithInstance(db, &migrateSql.Config{})
// 	if err != nil {
// 		panic("error while creating migration driver |  err: " + err.Error())
// 	}

// 	m, err := migrate.NewWithDatabaseInstance(os.Getenv("MIGRATION_SRC"), "mysql", migrationDriver)
// 	if err != nil {
// 		panic("error while creating migrate instance |  err: " + err.Error())
// 	}

// 	err = m.Up()
// 	if err != nil && err != migrate.ErrNoChange {
// 		panic("error while running migrations |  err: " + err.Error())
// 	}
// }

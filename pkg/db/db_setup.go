package db

import (
	"fmt"
	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
	"log"
)

func SetupDB() (*pg.DB, error) {
	// connect to db
	db := pg.Connect(&pg.Options{
		Addr:     "db:5432",
		User:     "postgres",
		Password: "postgres",
	})

	// run migrations
	collection := migrations.NewCollection()
	err := collection.DiscoverSQLMigrations("migrations")
	if err != nil {
		log.Println("DB error on DiscoverSQLMigrations")
		log.Printf("error nya adalah %s\n", err.Error())
		return nil, err
	}

	_, _, err = collection.Run(db, "init")
	if err != nil {
		log.Println("DB error on Run init")
		log.Printf("error nya adalah %s\n", err.Error())
		return nil, err
	}

	oldVersion, newVersion, err := collection.Run(db, "up")
	if err != nil {
		log.Println("error euy")
		return nil, err
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}

	// return the db connection
	return db, nil
}

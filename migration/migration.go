package migration

import (
	"fmt"

	"github.com/pilinux/postmark/model"

	"github.com/pilinux/gorest/config"
	"github.com/pilinux/gorest/database"

	"gorm.io/gorm"
)

// Load all the models
type postmarkOutbound model.PostmarkOutbound

var db *gorm.DB
var errorState int

// DBMigrate migrates the database
func DBMigrate(dropOldTables bool) {
	/*
	** 0 = default/no error
	** 1 = error
	**/
	errorState = 0

	db = database.InitDB()

	// Auto migration
	/*
		- Automatically migrate schema
		- Only create tables with missing columns and missing indexes
		- Will not change/delete any existing columns and their types
	*/

	// Careful! It will drop all the tables!
	if dropOldTables {
		DropAllTables()
	}

	// Automatically migrate all the tables
	MigrateTables()

	if errorState != 0 {
		fmt.Println("Auto migration failed!")
		return
	}

	fmt.Println("Auto migration is completed! Starting the application...")
}

// DropAllTables drop all the tables
func DropAllTables() {
	// Careful! It will drop all the tables!
	if err := db.Migrator().DropTable(&postmarkOutbound{}); err != nil {
		errorState = 1
		fmt.Println(err)
	} else {
		fmt.Println("Old tables are deleted!")
	}
}

// MigrateTables migrates the database - MySQL / PostgreSQL / SQLite3
func MigrateTables() {
	configureDB := config.Database().RDBMS
	driver := configureDB.Env.Driver

	if driver == "mysql" {
		// db.Set() --> add table suffix during auto migration
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
			&postmarkOutbound{},
		); err != nil {
			errorState = 1
			fmt.Println(err)
		} else {
			fmt.Println("New tables are  migrated successfully!")
		}
	} else {
		if err := db.AutoMigrate(
			&postmarkOutbound{},
		); err != nil {
			errorState = 1
			fmt.Println(err)
		} else {
			fmt.Println("New tables are  migrated successfully!")
		}
	}
}

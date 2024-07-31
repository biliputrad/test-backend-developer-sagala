package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"test-backend-developer-sagala/common/constants"
	"test-backend-developer-sagala/config/database"
	"test-backend-developer-sagala/config/env"
)

func ConfigurationPostgres(config env.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v",
		config.DbHost, config.DbUsername, config.DbPassword, config.DbName, config.DbPort, config.DbTz,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}, &gorm.Config{
		Logger: logger.Default.LogMode(database.DatabaseLogger(config.DbLogLevel)),
	})

	if err != nil {
		message := fmt.Sprintf("%s database connection failed", constants.Database)
		log.Fatal(message)
	}
	log.Printf("%s successfully connected to database %s", constants.Database, config.DbName)

	if config.DbMigrate {
		log.Printf("%s migrating tables...", constants.Database)
		err = database.MigrateTables(db)
		if err != nil {
			message := fmt.Sprintf("%s migrations failed", constants.Database)
			log.Fatal(message)
		}
		log.Printf("%s migrations success", constants.Database)
	}

	GlobalDatabase = db
	return GlobalDatabase
}

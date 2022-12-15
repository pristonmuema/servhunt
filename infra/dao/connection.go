package dao

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"servhunt/config"
	"servhunt/infra/utils"
)

var (
	logger = utils.GetRootLogger()
)

type Repository struct {
	DB *gorm.DB
}

func InitRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func Connection(config *config.Config) *gorm.DB {
	DbName := config.Database.Name
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", config.Database.User, config.Database.Password,
		config.Database.ConnectionUrl, DbName)
	repo, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		logger.Error("error connecting to database", zap.NamedError("error.message", err))
		return nil
	}

	err = repo.Exec("CREATE DATABASE IF NOT EXISTS " + DbName).Error
	if err != nil {
		logger.Error("error creating database")
	}
	db, er := repo.DB()
	if er != nil {
		logger.Error("error while accessing database")
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	erf := db.Ping()
	if erf != nil {
		logger.Error("error when pinging database")
	}
	logger.Info("Connection to the database " + DbName + " is successful ")

	err = repo.Exec("USE " + DbName).Error
	if err != nil {
		logger.Error("error selecting database")
	}

	return repo
}

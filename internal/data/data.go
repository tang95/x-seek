package data

import (
	"errors"
	"github.com/tang95/x-seek/config"
	"github.com/tang95/x-seek/internal/model"
	"github.com/tang95/x-seek/internal/service"
	"go.uber.org/zap"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Data struct {
	config       *config.Server
	logger       *zap.Logger
	database     *gorm.DB
	IncidentRepo model.IncidentRepo
	TeamRepo     model.TeamRepo
	UserRepo     model.UserRepo
	Transaction  service.Transaction
}

func NewData(config *config.Server, logger *zap.Logger) (*Data, error) {
	database, err := connectDatabase(config.Database.Driver, config.Database.Source)
	if err != nil {
		return nil, err
	}
	err = database.AutoMigrate(
		&model.Incident{},
		&model.IncidentUser{},
		&model.IncidentActivity[any]{},
		&model.User{},
		&model.Team{},
		&model.TeamUser{},
	)
	if err != nil {
		logger.Error("failed to auto migrate database", zap.Error(err))
		return nil, err
	}
	d := &Data{
		config:   config,
		logger:   logger,
		database: database,
	}
	d.Transaction = newTransaction(d)
	d.IncidentRepo = newIncidentRepo(d)
	d.TeamRepo = newTeamRepo(d)
	d.UserRepo = newUserRepo(d)
	return d, nil
}

func connectDatabase(driver, source string) (*gorm.DB, error) {
	var (
		database *gorm.DB
		err      error
	)
	switch driver {
	case "mysql":
		database, err = gorm.Open(mysql.Open(source), &gorm.Config{})
		if err != nil {
			return nil, err
		}
	case "sqlite":
		database, err = gorm.Open(sqlite.Open(source), &gorm.Config{})
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("unsupported database driver")
	}
	return database, nil
}

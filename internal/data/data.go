package data

import (
	"demo-service/internal/biz"
	"demo-service/internal/conf"
	"gorm.io/driver/sqlite"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDemoRepo)

// Data .
type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	var db *gorm.DB
	var err error
	if c.Database.Driver == "sqlite" {
		db, err = gorm.Open(sqlite.Open(c.Database.Source), &gorm.Config{})
		db.AutoMigrate(&biz.Demo{})
		if err != nil {
			return &Data{}, nil, err
		}
	}
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		dbs, _ := db.DB()
		dbs.Close()
	}
	return &Data{db: db}, cleanup, nil
}

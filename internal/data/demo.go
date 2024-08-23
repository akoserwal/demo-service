package data

import (
	"context"
	"demo-service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type DemoRepo struct {
	data *Data
	log  *log.Helper
}

// Save implements biz.DemoRepo.
func (d *DemoRepo) Save(ctx context.Context, model *biz.Demo) (*biz.Demo, error) {
	if err := d.data.db.Session(&gorm.Session{FullSaveAssociations: true}).Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (d *DemoRepo) List(ctx context.Context, count int64) (*[]biz.Demo, error) {
	var results []biz.Demo
	re := d.data.db.Limit(int(count)).Find(&results)
	d.log.Infof("fetch results: %v", results)
	if re.Error != nil {
		d.log.Error(re.Error)
		return nil, re.Error
	}
	return &results, nil
}

// NewdemoRepo .
func NewDemoRepo(data *Data, logger log.Logger) biz.DemoRepo {
	return &DemoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

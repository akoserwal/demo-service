package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Demo model
type Demo struct {
	Name  string
	Email string
}

// GreeterRepo is a Greater repo.
type DemoRepo interface {
	Save(context.Context, *Demo) (*Demo, error)
	List(ctx context.Context, count int64) (*[]Demo, error)
}

// DemoUsecase is a  Demo usecase.
type DemoUsecase struct {
	repo DemoRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Demo usecase.
func NewDemoUsecase(repo DemoRepo, logger log.Logger) *DemoUsecase {
	return &DemoUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Demo, and returns the new Demo.
func (uc *DemoUsecase) CreateDemo(ctx context.Context, g *Demo) (*Demo, error) {
	uc.log.WithContext(ctx).Infof("CreateDemo: %v", g.Name)
	return uc.repo.Save(ctx, g)
}

func (uc *DemoUsecase) ListDemos(ctx context.Context, count int64) (*[]Demo, error) {
	uc.log.WithContext(ctx).Infof("listDemo: %v", count)
	return uc.repo.List(ctx, count)
}

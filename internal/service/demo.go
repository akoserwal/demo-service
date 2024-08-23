package service

import (
	"context"
	"demo-service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	pb "demo-service/api/testdomain/v1"
)

type DemoService struct {
	pb.UnimplementedDemoserviceServer
	demouc *biz.DemoUsecase
	log    *log.Helper
}

func NewDemoService(demouc *biz.DemoUsecase, logger log.Logger) *DemoService {
	return &DemoService{demouc: demouc, log: log.NewHelper(logger)}
}

func (s *DemoService) CreateDemo(ctx context.Context, req *pb.CreateDemoRequest) (*pb.CreateDemoReply, error) {
	s.log.Info(req.Demo)
	r := biz.Demo{
		Name:  req.Demo.Name,
		Email: req.Demo.Email,
	}

	_, err := s.demouc.CreateDemo(ctx, &r)
	if err != nil {
		return nil, err
	}
	return &pb.CreateDemoReply{Message: "created"}, nil
}
func (s *DemoService) UpdateDemo(ctx context.Context, req *pb.UpdateDemoRequest) (*pb.UpdateDemoReply, error) {
	return &pb.UpdateDemoReply{}, nil
}
func (s *DemoService) DeleteDemo(ctx context.Context, req *pb.DeleteDemoRequest) (*pb.DeleteDemoReply, error) {
	return &pb.DeleteDemoReply{}, nil
}
func (s *DemoService) GetDemo(ctx context.Context, req *pb.GetDemoRequest) (*pb.GetDemoReply, error) {
	return &pb.GetDemoReply{}, nil
}
func (s *DemoService) ListDemo(ctx context.Context, req *pb.ListDemoRequest) (*pb.ListDemoReply, error) {
	demos, err := s.demouc.ListDemos(ctx, req.Count)
	if err != nil {
		return nil, err
	}
	rs := make([]*pb.Demo, len(*demos))
	for _, x := range *demos {
		rs = append(rs, &pb.Demo{
			Name:  x.Name,
			Email: x.Email,
		})
	}
	return &pb.ListDemoReply{Demo: rs}, nil
}

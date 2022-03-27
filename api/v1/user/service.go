package user

import "github.com/harisaginting/ginting/pkg/tracer"
import srv "github.com/harisaginting/ginting/service/samplegrpc"

type Service struct {
	repo Repository
}

func ProviderService(r Repository) Service {
	return Service{
		repo: r,
	}
}

func (service *Service) List(res *ResponseList) {
	trace := tracer.Span("ListUser")
	defer trace.End()

	users := service.repo.FindAll()
	res.Items = users
	res.Total = len(users)

	tracer.SetAttributeInt(trace,"total User",res.Total)
	return
}

func (service *Service) ListGRPC(param string) (res string) {
	trace := tracer.Span("ListUser")
	defer trace.End()

	// CALL GRPC SAMPLE
	res = srv.Sample("TEST SERVICE LIST GRPC")
	// END CALL GRPC SAMPLE

	tracer.SetAttributeInt(trace,"total User",res)
	return
}
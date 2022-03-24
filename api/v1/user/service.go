package user

import "github.com/harisaginting/ginting/pkg/tracer"

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
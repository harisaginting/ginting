package user

type Service struct {
	repo Repository
}

func ProviderService(r Repository) Service {
	return Service{
		repo: r,
	}
}

func (service *Service) List(res *ResponseList) {
	users := service.repo.FindAll()
	res.Items = users
	res.Total = len(users)
	return
}
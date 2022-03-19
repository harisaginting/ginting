package user


type Service struct {
	repository Repository
}

func ProviderService(r Repository) Service {
	return Service{
		repository: r,
	}
}
package services

type uservice struct {
}

type UserService interface {
	HandleGetUser(string) string
}

func NewService() UserService {
	return &uservice{}
}

func (u *uservice) HandleGetUser(id string) string {
	return "user1"
}

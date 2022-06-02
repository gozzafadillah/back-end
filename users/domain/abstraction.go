package domain_users

type Service interface {
	Login(email, password string) (string, error)

	Register(domain Users) (Users, error)
}

type Repository interface {
	CheckEmailPassword(email, password string) (Users, error)

	Store(domain Users) (int, error)
	GetById(id int) (Users, error)
}

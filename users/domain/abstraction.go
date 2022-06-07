package domain_users

type Service interface {
	// Users
	Login(email, password string) (string, error)
	Register(domain Users) (string, error)
	GetUserPhone(phone string) (Users, error)
	// Account
	InsertAccount(domain Account) (Account, error)
	// Address
}

type Repository interface {
	// Users
	CheckEmailPassword(email, password string) (Users, error)
	Store(domain Users) (int, error)
	GetById(id int) (Users, error)
	GetByPhone(phone string) (Users, error)
	// Account
	StoreAccount(domain Account) (Account, error)
}

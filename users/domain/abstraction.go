package domain_users

type Service interface {
	// Users
	Login(email, password string) (string, error)
	Register(domain Users) (Users, error)
	GetUsers() ([]Users, error)
	GetUserPhone(phone string) (Users, error)
	EditUser(phone string, domain Users) error
	// Account
	InsertAccount(domain Account) (Account, error)
	GetUserAccount(phone string) (Account, error)
	// Verif User
	Verif(code string) (string, error)
}

type Repository interface {
	// Users
	GetAllUser() ([]Users, error)
	CheckEmailPassword(email, password string) (Users, error)
	Store(domain Users) (string, error)
	GetByPhone(phone string) (Users, error)
	Update(phone string, domain Users) error
	// Account
	StoreAccount(domain Account) (Account, error)
	GetUserAccount(phone string) (Account, error)
	// Users Verification
	StoreVerif(domain UserVerif) (string, error)
	CheckOTP(phone string) (bool, error)
}

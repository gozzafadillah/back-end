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
	GetUserAccount(phone string) Account
	// Verif User
	AddUserVerif(code, email, name string) error
	Verif(code string) error
	// admin-dashboard
	CountUsersCustomer() int
}

type Repository interface {
	// Users
	GetAllUser() ([]Users, error)
	CheckEmailPassword(email, password string) (Users, error)
	Store(domain Users) (string, error)
	GetByPhone(phone string) (Users, error)
	ChangeStatusUsers(email string) error
	Update(phone string, domain Users) error
	// Account
	StoreAccount(domain Account) (Account, error)
	GetUserAccount(phone string) (Account, error)
	// Users Verification
	StoreOtpUserVerif(code string, email string) error
	Verif(code string) (UserVerif, error)
	ChangeStatusVerif(email string) error
	// admin-dashboard
	Count() int
}

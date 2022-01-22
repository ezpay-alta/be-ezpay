package users

type Core struct {
	ID       int
	Fullname string
	Phone    string
	Email    string
	Role     string
	Password string
}

type Business interface {
	CreateUser(data Core) (userId int, err error)
	GetAllUsers() ([]Core, error)
	GetUserById(id int) (Core, error)
	UpdateUserById(userId int, data Core) error
	DeleteUserById(userId int) error
}

type Data interface {
	CreateUser(data Core) (userId int, err error)
	GetAllUsers() ([]Core, error)
	GetUserById(userId int) (Core, error)
	UpdateUserById(userId int, data Core) error
	DeleteUserById(userId int) error
}

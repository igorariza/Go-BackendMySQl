package users

import (
	"database/sql"

	users "github.com/igorariza/Go-BackendMySQl/users/models"
)

type UserGateway interface {
	CreateUser(p *users.CreateUserCMD) (*users.User, error)
	GetUsers() []*users.User
	GetUserByID(id int64) (*users.User, error)
	GetUserByEmail(email string) (*users.User, error)
}
type CreateUserInDB struct {
	UserStorage
}

func (c *CreateUserInDB) CreateUser(p *users.CreateUserCMD) (*users.User, error) {
	return c.createUserDB(p)
}

func (c *CreateUserInDB) GetUsers() []*users.User {
	return c.getUsersDB()
}

func (c *CreateUserInDB) GetUserByID(id int64) (*users.User, error) {
	return c.getUserByIDBD(id)
}

func (c *CreateUserInDB) GetUserByEmail(email string) (*users.User, error) {
	return c.getUserByEmailBD(email)
}

func NewUserGateway(db *sql.DB) UserGateway {
	return &CreateUserInDB{NewUserStorageGateway(db)}
}

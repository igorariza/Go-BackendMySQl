package users

import (
	"database/sql"

	users "github.com/igorariza/Go-BackendMySQl/users/models"
)

//UserGateway comment gteneric
type UserGateway interface {
	CreateUser(p *users.CreateUserCMD) (*users.User, error)
	GetUsers() []*users.User
	GetUserByID(id int64) (*users.User, error)
	GetUserByEmail(email string) (*users.User, error)
}

//CreateUserInDB comment gteneric
type CreateUserInDB struct {
	UserStorage
}

//CreateUser comment gteneric
func (c *CreateUserInDB) CreateUser(p *users.CreateUserCMD) (*users.User, error) {
	return c.createUserDB(p)
}

//GetUsers comment gteneric
func (c *CreateUserInDB) GetUsers() []*users.User {
	return c.getUsersDB()
}

//GetUserByID comment gteneric
func (c *CreateUserInDB) GetUserByID(id int64) (*users.User, error) {
	return c.getUserByIDBD(id)
}

//GetUserByEmail comment gteneric
func (c *CreateUserInDB) GetUserByEmail(email string) (*users.User, error) {
	return c.getUserByEmailBD(email)
}

//NewUserGateway comment gteneric
func NewUserGateway(db *sql.DB) UserGateway {
	return &CreateUserInDB{NewUserStorageGateway(db)}
}

package users

import (
	"database/sql"

	users "github.com/igorariza/Go-BackendMySQl/users/models"
)

//LoginGateway comment generic
type LoginGateway interface {
	LoginUser(p *users.LoginUser) (*users.User, error)
}

//LoginInDB comment generic
type LoginInDB struct {
	UserStorage
}

//LoginUser comment generic
func (c *LoginInDB) LoginUser(p *users.LoginUser) (*users.User, error) {
	return c.loginUserDB(p)
}

//NewLoginUserGateway comment generic
func NewLoginUserGateway(db *sql.DB) LoginGateway {
	return &LoginInDB{NewUserLoginStorageGateway(db)}
}

package users

import (
	"database/sql"
	"log"
	"time"

	"github.com/igorariza/Go-BackendMySQl/internal/storage"
	users "github.com/igorariza/Go-BackendMySQl/users/models"
	"golang.org/x/crypto/bcrypt"
)

//UserStorage interface rutas
type UserStorage interface {
	createUserDB(p *users.CreateUserCMD) (*users.User, error)
	getUsersDB() []*users.User
	getUserByIDBD(id int64) (*users.User, error)
	getUserByEmailBD(email string) (*users.User, error)

	loginUserDB(p *users.LoginUser) (*users.User, error)
}

//UserService db mysql sql.DB
type UserService struct {
	db *sql.DB
}

//NewUserStorageGateway comment generic
func NewUserStorageGateway(db *sql.DB) UserStorage {
	return &UserService{db: db}
}

//NewUserLoginStorageGateway comment generic
func NewUserLoginStorageGateway(db *sql.DB) UserStorage {
	return &UserService{db: db}
}

//loginUserDB comment generic
func (s *UserService) loginUserDB(p *users.LoginUser) (*users.User, error) {
	var user users.User
	// p.Password, _ = storage.EncryptPassword(p.Password)
	var passwordUser string
	query := "SELECT idUser, document_id, first_name, last_name, email, password, phone, address, photo, created_at, type_id, date_birth, rh, idSede, is_active FROM Users WHERE email = ?"
	err := s.db.QueryRow(query, p.Email).Scan(&user.ID, &user.DocumentID, &user.FirstName, &user.LastName, &user.Email, &passwordUser, &user.Phone, &user.Address, &user.Photo, &user.CreatedAt, &user.TypeID, &user.DateBirth, &user.Rh, &user.IDSede, &user.IsActive)

	if err != nil {
		log.Printf("cannot fetch user email ")
		return nil, err
	}

	passwordBytes := []byte(p.Password)
	passwordBD := []byte(passwordUser)

	err = bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		log.Printf("error en la contraseña ")
		return nil, nil
	}
	return &user, nil
}

//createUserDB comment generic
func (s *UserService) createUserDB(p *users.CreateUserCMD) (*users.User, error) {
	p.Password, _ = storage.EncryptPassword(p.Password)
	p.CreatedAt = time.Now().String()

	res, err := s.db.Exec("insert into Users (document_id, first_name, last_name, email, password, phone, address, photo, created_at, type_id, date_birth, rh, idSede, is_active) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		p.DocumentID, p.FirstName, p.LastName, p.Email, p.Password, p.Phone, p.Address, p.Photo, p.CreatedAt, p.TypeID, p.DateBirth, p.Rh, p.IDSede, p.IsActive)

	if err != nil {
		log.Printf("cannot save the user, %s", err.Error())
		return nil, err
	}

	id, err := res.LastInsertId()

	return &users.User{
		ID:         id,
		DocumentID: p.DocumentID,
		FirstName:  p.FirstName,
		LastName:   p.LastName,
		Email:      p.Email,
		Phone:      p.Phone,
		Address:    p.Address,
		Photo:      p.Photo,
		CreatedAt:  p.CreatedAt,
		TypeID:     p.TypeID,
		DateBirth:  p.DateBirth,
		LastAccess: p.LastAccess,
		Rh:         p.Rh,
		IDSede:     p.IDSede,
		IsActive:   p.IsActive,
	}, nil
}

//getUsersDB comment generic
func (s *UserService) getUsersDB() []*users.User {
	rows, err := s.db.Query("select _id, created, name, last_name, email, from Users")

	if err != nil {
		log.Printf("cannot execute select query: %s", err.Error())
		return nil
	}
	defer rows.Close()
	var p []*users.User
	for rows.Next() {
		var user users.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Address, &user.Phone,
			&user.Email, &user.CreatedAt)
		if err != nil {
			log.Println("cannot read current row ", err)
			return nil
		}
		p = append(p, &user)
	}

	return p
}

//getUserByEmailBD comment generic
func (s *UserService) getUserByIDBD(id int64) (*users.User, error) {
	var user users.User
	err := s.db.QueryRow(`select id, first_name, last_name, address, phone, email, created_at from users
		where id = ?`, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Address, &user.Phone,
		&user.Email, &user.CreatedAt)

	if err != nil {
		log.Printf("cannot fetch user %v", err)
		return nil, err
	}

	return &user, nil
}

//getUserByEmailBD comment generic
func (s *UserService) getUserByEmailBD(email string) (*users.User, error) {
	var user users.User
	var emailUser string
	query := "SELECT email FROM Users WHERE email = ?"
	err := s.db.QueryRow(query, email).Scan(&emailUser)

	if err != nil {
		log.Printf("cannot fetch user email " + err.Error())
		return nil, err
	}
	user.Email = emailUser

	return &user, nil
}

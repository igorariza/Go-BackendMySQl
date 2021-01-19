package users

import (
	"database/sql"
	"log"
	"time"

	"github.com/igorariza/Go-BackendMySQl/internal/storage"
	users "github.com/igorariza/Go-BackendMySQl/users/models"
)

//UserStorage interface rutas
type UserStorage interface {
	createUserDB(p *users.CreateUserCMD) (*users.User, error)
	getUsersDB() []*users.User
	getUserByIDBD(id int64) (*users.User, error)
	getUserByEmailBD(email string) (*users.User, error)
}

//UserService db mysql sql.DB
type UserService struct {
	db *sql.DB
}

func NewUserStorageGateway(db *sql.DB) UserStorage {
	return &UserService{db: db}
}

func (s *UserService) createUserDB(p *users.CreateUserCMD) (*users.User, error) {
	p.Password, _ = storage.EncryptPassword(p.Password)
	p.CreatedAt = time.Now().String()

	res, err := s.db.Exec("insert into Users (document_id, first_name, last_name, email, password, phone, address, photo, created_at, type_id, date_birth, rh, idSede, is_active) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		p.DocumentId, p.FirstName, p.LastName, p.Email, p.Password, p.Phone, p.Address, p.Photo, p.CreatedAt, p.TypeId, p.DateBirth, p.Rh, p.IdSede, p.IsActive)

	if err != nil {
		log.Printf("cannot save the user, %s", err.Error())
		return nil, err
	}

	id, err := res.LastInsertId()

	return &users.User{
		ID:         id,
		DocumentId: p.DocumentId,
		FirstName:  p.FirstName,
		LastName:   p.LastName,
		Email:      p.Email,
		Phone:      p.Phone,
		Address:    p.Address,
		Photo:      p.Photo,
		CreatedAt:  p.CreatedAt,
		TypeId:     p.TypeId,
		DateBirth:  p.DateBirth,
		LastAccess: p.LastAccess,
		Rh:         p.Rh,
		IdSede:     p.IdSede,
		IsActive:   p.IsActive,
	}, nil
}

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

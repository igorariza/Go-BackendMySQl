package users

import (
	"database/sql"
	"log"
	"time"

	sto "github.com/igorariza/Go-BackendMySQl/internal/storage"
	usr "github.com/igorariza/Go-BackendMySQl/users/models"
	"golang.org/x/crypto/bcrypt"
)

//UserStorage interface rutas
type UserStorage interface {
	createUserDB(p *usr.CreateUserCMD) (*usr.User, error)
	getUsersDB() []*usr.User
	getUserByIDBD(id int64) (*usr.User, error)
	getUserByEmailBD(email string) (*usr.User, error)
	loginUserDB(p *usr.LoginUser) (*usr.User, error)
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
func (s *UserService) loginUserDB(p *usr.LoginUser) (*usr.User, error) {
	var user usr.User
	var passwordUser string
	query := "SELECT users.idUser, users.document_id, users.first_name, users.last_name, users.email, users.password, users.phone, users.address, users.photo, users.created_at, users.type_id, users.date_birth, users.rh, users.idSede, users.is_active, sede.name_sede FROM users INNER JOIN sede ON sede.idSede = users.idSede WHERE email = ?"
	err := s.db.QueryRow(query, p.Email).Scan(&user.ID, &user.DocumentID, &user.FirstName, &user.LastName, &user.Email, &passwordUser, &user.Phone, &user.Address, &user.Photo, &user.CreatedAt, &user.TypeID, &user.DateBirth, &user.Rh, &user.IDSede, &user.IsActive, &user.NameSede)

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
func (s *UserService) createUserDB(p *usr.CreateUserCMD) (*usr.User, error) {

	//existe, _, err := sto.ChequeoYaExisteUsuario(p.DocumentID)

	// if existe != 0 {
	// 	log.Printf("Usuario ya existe, " + strconv.Itoa(existe))
	// 	return nil, err
	// }
	//log.Printf("Usuario ya existe, " + strconv.Itoa(existe))
	p.Password, _ = sto.EncryptPassword(p.Password)
	p.CreatedAt = time.Now().String()

	res, err := s.db.Exec("insert into users (document_id, first_name, last_name, email, password, phone, address, photo, created_at, type_id, date_birth, last_access, rh, idSede, is_active) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		p.DocumentID, p.FirstName, p.LastName, p.Email, p.Password, p.Phone, p.Address, p.Photo, p.CreatedAt, p.TypeID, p.DateBirth, p.LastAccess, p.Rh, p.IDSede, p.IsActive)

	if err != nil {
		log.Printf("cannot save the user, %s", err.Error())
		return nil, err
	}

	id, err := res.LastInsertId()

	return &usr.User{
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
func (s *UserService) getUsersDB() []*usr.User {
	rows, err := s.db.Query("select _id, created, name, last_name, email, from Users")

	if err != nil {
		log.Printf("cannot execute select query: %s", err.Error())
		return nil
	}
	defer rows.Close()
	var p []*usr.User
	for rows.Next() {
		var user usr.User
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
func (s *UserService) getUserByIDBD(id int64) (*usr.User, error) {
	var user usr.User
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
func (s *UserService) getUserByEmailBD(email string) (*usr.User, error) {
	var user usr.User
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

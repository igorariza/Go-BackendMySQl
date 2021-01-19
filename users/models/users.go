package users

type User struct {
	ID         int64  `json:"idUser"`
	DocumentId string `json:"document_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	Photo      string `json:"photo"`
	CreatedAt  string `json:"created_at"`
	TypeId     string `json:"type_id"`
	DateBirth  string `json:"date_birth"`
	LastAccess string `json:"last_access"`
	Rh         string `json:"rh"`
	IdSede     string `json:"idSede"`
	IsActive   string `json:"is_active"`
}

type CreateUserCMD struct {
	ID         int64  `json:"idUser"`
	DocumentId string `json:"document_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	Photo      string `json:"photo"`
	CreatedAt  string `json:"created_at"`
	TypeId     string `json:"type_id"`
	DateBirth  string `json:"date_birth"`
	LastAccess string `json:"last_access"`
	Rh         string `json:"rh"`
	IdSede     string `json:"idSede"`
	IsActive   string `json:"is_active"`
}

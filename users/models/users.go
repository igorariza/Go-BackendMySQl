package users

//User comment generic
type User struct {
	ID         int64  `json:"idUser"`
	DocumentID string `json:"document_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	Photo      string `json:"photo"`
	CreatedAt  string `json:"created_at"`
	TypeID     string `json:"type_id"`
	DateBirth  string `json:"date_birth"`
	LastAccess string `json:"last_access"`
	Rh         string `json:"rh"`
	IDSede     string `json:"idSede"`
	IsActive   string `json:"is_active"`
	NameSede   string `json:"name_sede"`
}

//CreateUserCMD comment generic
type CreateUserCMD struct {
	ID         int64  `json:"idUser"`
	DocumentID string `json:"document_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	Photo      string `json:"photo"`
	CreatedAt  string `json:"created_at"`
	TypeID     string `json:"type_id"`
	DateBirth  string `json:"date_birth"`
	LastAccess string `json:"last_access"`
	Rh         string `json:"rh"`
	IDSede     string `json:"idSede"`
	IsActive   string `json:"is_active"`
}

//LoginUser comment generic
type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

package users

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/igorariza/Go-BackendMySQl/internal/web"
	users "github.com/igorariza/Go-BackendMySQl/users/gateway"
	models "github.com/igorariza/Go-BackendMySQl/users/models"
)

//UserCreateHTTPService comment login
type UserCreateHTTPService struct {
	gtw users.UserGateway
}

//NewUserCreateHTTPService comment login
func NewUserCreateHTTPService(db *sql.DB) *UserCreateHTTPService {
	return &UserCreateHTTPService{
		users.NewUserGateway(db),
	}
}

//GetUsersHandler comment login
func (s *UserCreateHTTPService) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	p := s.gtw.GetUsers()
	if p == nil || len(p) == 0 {
		p = []*models.User{}
	}
	web.Success(&p, http.StatusOK).Send(w)
}

//GetUsersByIDHandler comment login
func (s *UserCreateHTTPService) GetUsersByIDHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	id, _ := strconv.ParseInt(userID, 10, 64)
	user, err := s.gtw.GetUserByID(id)

	if err != nil {
		web.ErrBadRequest.Send(w)
		return
	}

	web.Success(&user, http.StatusOK).Send(w)
}

//GetUsersByEmailHandler comment login
func (s *UserCreateHTTPService) GetUsersByEmailHandler(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	_, err := s.gtw.GetUserByEmail(email)

	if err != nil {
		web.Duplicate(nil, http.StatusBadRequest).Send(w)
		return
	}

	web.Success(nil, http.StatusOK).Send(w)
}

//CreateUsersHandler comment login
func (s *UserCreateHTTPService) CreateUsersHandler(w http.ResponseWriter, r *http.Request) {
	var cmd models.CreateUserCMD
	body := r.Body
	defer body.Close()
	err := json.NewDecoder(body).Decode(&cmd)

	if err != nil {
		web.ErrInvalidJSON.Send(w)
		log.Println("ErrInvalidJSON " + err.Error())
		return
	}

	user, err := s.gtw.CreateUser(&cmd)

	if err != nil {
		web.ErrBadRequest.Send(w)
		log.Println("ErrBadRequest " + err.Error())
		return
	}

	web.Success(&user, http.StatusOK).Send(w)
}

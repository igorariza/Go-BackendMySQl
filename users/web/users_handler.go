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

type UserHTTPService struct {
	gtw users.UserGateway
}

func NewUserHTTPService(db *sql.DB) *UserHTTPService {
	return &UserHTTPService{
		users.NewUserGateway(db),
	}
}

func (s *UserHTTPService) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	p := s.gtw.GetUsers()
	if p == nil || len(p) == 0 {
		p = []*models.User{}
	}
	web.Success(&p, http.StatusOK).Send(w)
}

func (s *UserHTTPService) GetUsersByIDHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	id, _ := strconv.ParseInt(userID, 10, 64)
	user, err := s.gtw.GetUserByID(id)

	if err != nil {
		web.ErrBadRequest.Send(w)
		return
	}

	web.Success(&user, http.StatusOK).Send(w)
}

func (s *UserHTTPService) GetUsersByEmailHandler(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	_, err := s.gtw.GetUserByEmail(email)

	if err != nil {
		web.Duplicate(nil, http.StatusBadRequest).Send(w)
		return
	}

	web.Success(nil, http.StatusOK).Send(w)
}

func (s *UserHTTPService) CreateUsersHandler(w http.ResponseWriter, r *http.Request) {
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

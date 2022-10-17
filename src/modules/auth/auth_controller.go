package auth

import (
	"encoding/json"
	"net/http"

	"github.com/adiet95/go-order-api/src/database/models"
	"github.com/adiet95/go-order-api/src/interfaces"
	"github.com/adiet95/go-order-api/src/libs"
)

type user_ctrl struct {
	repo interfaces.AuthService
}

func NewCtrl(reps interfaces.AuthService) *user_ctrl {
	return &user_ctrl{reps}
}

func (u *user_ctrl) SignIn(w http.ResponseWriter, r *http.Request) {
	var data models.User

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		libs.New(err.Error(), 401, true)
		return
	}

	u.repo.Login(data, w).Send(w)
}

func (u *user_ctrl) Register(w http.ResponseWriter, r *http.Request) {
	var data *models.User

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		libs.New(err.Error(), 401, true)
		return
	}
	u.repo.Register(data).Send(w)
}
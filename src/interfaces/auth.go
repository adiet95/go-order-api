package interfaces

import (
	"net/http"

	"github.com/adiet95/go-order-api/src/database/models"
	"github.com/adiet95/go-order-api/src/libs"
)

type AuthRepo interface {
	FindByEmail(username string) (*models.User, error)
	RegisterEmail(data *models.User) (*models.User, error)
}

type AuthService interface {
	Login(body models.User, w http.ResponseWriter) *libs.Response
	Register(body *models.User) *libs.Response
}
